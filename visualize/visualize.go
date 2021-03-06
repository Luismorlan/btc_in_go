package visualize

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/Luismorlan/btc_in_go/model"
	"github.com/Luismorlan/btc_in_go/utils"
	"github.com/Luismorlan/btc_in_go/visualize/memviz"
)

// We need to re-define the visualize model here because protobuf
// generated model contains lots of extra and unnecessary imformation
// that we don't really care.
type input struct {
	prevTxHash string
	index      int64
}

type output struct {
	value     float64
	publicKey string
}

type coinbaseTransaction struct {
	hash    string
	outputs []output
	height  int64
}

type transaction struct {
	hash    string
	inputs  []input
	outputs []output
}

type block struct {
	hash     string
	prevHash string
	coinbase coinbaseTransaction
	txs      []transaction
	nounce   int64
	children []block
}

// A fullnode in the chain
type Address struct {
	Ip   string
	Port string
}

type Node struct {
	EndPoint string
	Peers    []*Node
}

type Graph struct {
	Nodes map[Address]*Node
	Root  *Node
}

func NewAddress(ip string, port string) Address {
	return Address{Ip: ip, Port: port}
}

func NewNode(ip string, port string) *Node {
	return &Node{
		EndPoint: ip + ":" + port,
	}
}

func (n *Node) GetEndpoint() (string, string) {
	addr := EndpointToAddress(n.EndPoint)
	return addr.Ip, addr.Port
}

func (n *Node) AddPeer(peer *Node) {
	n.Peers = append(n.Peers, peer)
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[Address]*Node),
		Root:  nil,
	}
}

func EndpointToAddress(s string) Address {
	ss := strings.Split(s, ":")
	return Address{
		Ip:   ss[0],
		Port: ss[1],
	}
}

// Whether a node has already been in the graph.
func (g *Graph) HasNode(n *Node) bool {
	addr := EndpointToAddress(n.EndPoint)
	_, exist := g.Nodes[addr]
	return exist
}

// Add a single node to this graph.
func (g *Graph) AddNode(n *Node) {
	if g.HasNode(n) {
		return
	}
	addr := EndpointToAddress(n.EndPoint)
	g.Nodes[addr] = n
	// Add first encountered non nil node as root node.
	if g.Root == nil {
		g.Root = n
	}
}

// Get node, if not find create the node.
func (g *Graph) GetNode(addr Address) *Node {
	if _, exist := g.Nodes[addr]; !exist {
		node := NewNode(addr.Ip, addr.Port)
		g.Nodes[addr] = node
	}
	if g.Root == nil {
		g.Root = g.Nodes[addr]
	}
	return g.Nodes[addr]
}

// Given a tail node, return from chain of the d-th block to the tail,
// including the branch.
func constructData(tail *model.BlockWrapper, d int) block {
	r := tail
	// go to d blocks ago to find the root.
	for i := 0; i < d; i++ {
		if r.Parent == nil {
			break
		}
		r = r.Parent
	}

	// build the tree recursively.
	t := buildTree(r)
	return t
}

// The string of public key and hash is just too long to render, instead we take only first 3 and last 3
// characters and replace the middle part with '...'. E.g. "abcdefghi" will be rendered as "abc...ghi"
func shortenString(s string) string {
	if len(s) < 9 {
		return s
	}
	return fmt.Sprintf("%s...%s", s[0:3], s[len(s)-3:])
}

func shortenPK(s string) string {
	if len(s) < 9 {
		return s
	}
	mid := len(s) / 2
	i := mid - 1
	j := mid + 2
	return fmt.Sprintf("...%s...", s[i:j])
}

func txToTx(tx *model.Transaction) transaction {
	t := transaction{
		hash: shortenString(tx.Hash),
	}

	for i := 0; i < len(tx.Inputs); i++ {
		in := tx.Inputs[i]
		t.inputs = append(t.inputs, input{prevTxHash: in.PrevTxHash, index: in.Index})
	}

	for i := 0; i < len(tx.Outputs); i++ {
		out := tx.Outputs[i]
		t.outputs = append(t.outputs, output{publicKey: shortenPK(utils.BytesToHex(out.PublicKey)), value: out.Value})
	}
	return t
}

func txToCb(tx *model.Transaction) coinbaseTransaction {
	cb := coinbaseTransaction{
		hash: shortenString(tx.Hash),
	}

	for i := 0; i < len(tx.Outputs); i++ {
		out := tx.Outputs[i]
		cb.outputs = append(cb.outputs, output{publicKey: shortenPK(utils.BytesToHex(out.PublicKey)), value: out.Value})
	}

	cb.height = tx.Height
	return cb
}

func blockToblock(b *model.Block) block {
	n := block{
		hash:     shortenString(b.Hash),
		prevHash: shortenString(b.PrevHash),
		nounce:   b.Nounce,
	}

	if b.Coinbase != nil {
		n.coinbase = txToCb(b.Coinbase)
	}

	for i := 0; i < len(b.Txs); i++ {
		tx := b.Txs[i]
		n.txs = append(n.txs, txToTx(tx))
	}
	return n
}

// Recursively build the tree in a dfs manner.
func buildTree(root *model.BlockWrapper) block {
	node := blockToblock(root.B)
	for i := 0; i < len(root.Children); i++ {
		child := root.Children[i]
		childNode := buildTree(child)
		node.children = append(node.children, childNode)
	}
	return node
}

// Entry to this package, where:
// tail: tail of the entire blockchain as tracked by full node.
// d: depth to return.
// id: unique id of the full node.
func RenderBlockChain(tail *model.BlockWrapper, d int, id string) {
	buf := &bytes.Buffer{}

	chain := constructData(tail, d)

	memviz.Map(buf, &chain)

	// Write the parsed data to disk
	fileName := "/tmp/chaindata-" + id
	outputName := "/tmp/rendered-chain-" + id + ".png"
	err := ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("dot", "-Tpng", fileName, "-o", outputName)
	cmd.Run()

	opCmd := exec.Command("open", outputName)
	opCmd.Run()
}

func RenderGraph(g *Graph) {
	buf := &bytes.Buffer{}

	memviz.Map(buf, g.Root)

	id := "graph"
	// Write the parsed data to disk
	fileName := "/tmp/g-" + id
	outputName := "/tmp/rendered-g-" + id + ".png"
	err := ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("sfdp", "-Tpng", fileName, "-o", outputName)
	cmd.Run()

	opCmd := exec.Command("open", outputName)
	opCmd.Run()
}
