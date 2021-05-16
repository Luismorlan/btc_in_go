package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Luismorlan/btc_in_go/commands"
	"github.com/Luismorlan/btc_in_go/wallet"
)

const LOAD_OP string = "load"
const SEND_OP string = "send"
const DEFAULT_FN_IP string = "127.0.0.1"
const DEFAULT_FN_PORT string = "8280"

var (
	fPathFlag  *string
	ipAddrFlag *string
	portFlag   *string
)

func init() {
	fPathFlag = flag.String("fPath", "", "RSA file path for your private key")
	ipAddrFlag = flag.String("ip", DEFAULT_FN_IP, "ip address of full node")
	portFlag = flag.String("port", DEFAULT_FN_PORT, "Port number for connection")
}

func main() {
	flag.Parse()
	fmt.Println("fPath is", *fPathFlag)
	fmt.Println("ipAddr is", *ipAddrFlag)
	fmt.Println("port is", *portFlag)

	wallet, err := parseToWallet(*fPathFlag, *ipAddrFlag, *portFlag)
	if err != nil {
		log.Println("[ERROR]Failed to get your wallet. Please check the input.")
		return
	}
	log.Println("Wallet logged in")
	cmd := make(chan commands.ClientCommand)
	go ParseCommand(cmd)
	go HandleCommand(cmd, wallet)

	c := make(chan int)
	<-c
}

func ParseCommand(cmd chan commands.ClientCommand) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		c, err := commands.CreateClientCommand(text)
		if err != nil {
			log.Println(err)
			cmd <- c
		}
	}
}

func HandleCommand(cmd chan commands.ClientCommand, wallet *wallet.Wallet) {
	for {
		c := <-cmd
		switch c.Op {
		case commands.TRANSFER:
			log.Println("Money transfer initiated")
		case commands.MY_PUBLIC_KEY:
			fmt.Println(&wallet.Keys.PublicKey)
		default:
			log.Println("Unhandled command:", c)
			fmt.Print("> ")
		}
	}
}

func parseToWallet(fPath string, ipAddr string, port string) (*wallet.Wallet, error) {
	var wallet wallet.Wallet
	var err error
	wallet.FullNodeIp = ipAddr
	wallet.FullNodePort = port
	wallet.Keys, err = parseKeyFile(fPath)
	return &wallet, err
}

func parseKeyFile(fPath string) (*rsa.PrivateKey, error) {
	var userKey *rsa.PrivateKey
	var err error
	if fPath == "" || fPath == "\n" {
		log.Println("No key file detected. Will generate a new key for you")
		userKey, err = rsa.GenerateKey(rand.Reader, 2048)
		// TODO : add new key in user path
		if err != nil {
			log.Fatal("Got error when generating new key", err)
			return nil, err
		}
		return userKey, nil
	}
	// read from rsa file
	userKey, err = readKeyFromFPath(fPath)
	if err != nil {
		log.Fatal("Failed to read your key from path {} with error {}", fPath, err)
		return nil, err
	}
	return userKey, nil
}

func readKeyFromFPath(fPath string) (*rsa.PrivateKey, error) {
	fileContent, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(fileContent))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key, nil
}
