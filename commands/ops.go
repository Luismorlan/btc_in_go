package commands

import (
	"errors"
	"regexp"
	"strings"
)

type Operation int

const IP_REGEX = "(?:[0-9]{1,3}\\.){3}[0-9]{1,3}"
const PORT_REGEX = "[0-9]{4,5}"

const (
	DEFAULT = iota
	// Start mining, infinite loop until explicit cancel.
	START
	// Restart mining when new tail replace the tail we mine on.
	RESTART
	// Stop mining completely.
	STOP
	// Add a new peer to this full node.
	ADD_PEER
	// Renove a peer by ip and port.
	REMOVE_PEER
)

// A command contains a operation and many arguments.
type Command struct {
	Op   Operation
	Args []string
}

func (c Command) IsValid() bool {
	switch c.Op {
	case START, RESTART, STOP:
		return len(c.Args) == 0
	case ADD_PEER, REMOVE_PEER:
		if len(c.Args) != 2 {
			return false
		}
		ipAddr := c.Args[0]
		port := c.Args[1]
		ipRegex, _ := regexp.Compile(IP_REGEX)
		portRegex, _ := regexp.Compile(PORT_REGEX)
		return ipRegex.Match([]byte(ipAddr)) && portRegex.Match([]byte(port))
	default:
		return false
	}
}

// From string, create
func CreateCommand(s string) (Command, error) {
	// split command by space.
	ss := strings.Split(s, " ")
	if len(ss) == 0 {
		return Command{}, errors.New("command is empty")
	}
	cmd := Command{}
	switch ss[0] {
	case "start":
		cmd.Op = START
	case "restart":
		cmd.Op = RESTART
	case "stop":
		cmd.Op = STOP
	case "add_peer":
		cmd.Op = ADD_PEER
	case "remove_peer":
		cmd.Op = REMOVE_PEER
	}
	cmd.Args = ss[1:]
	if !cmd.IsValid() {
		return Command{}, errors.New("invalid command")
	}
	return cmd, nil
}

// Create a brand new command with default operation.
func NewDefaultCommand() Command {
	return Command{
		Op: DEFAULT,
	}
}

func (c Command) IsDefault() bool {
	return c.Op == DEFAULT
}
