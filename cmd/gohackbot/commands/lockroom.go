package commands

import (
	"github.com/0x17de/gohackchat/pkg/hack"
)

type LockRoomCommand struct {
	Command
}

func NewLockRoomCommand() *LockRoomCommand {
	return &LockRoomCommand{Command{
		aliases:     []string{"lockroom", "unlockroom"},
		description: "(Un-)Lock the current room",
	}}
}

func (cmd *LockRoomCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)

	reply := make(hack.JsonValue)
	if args[0] == "lockroom" {
		reply["cmd"] = "lockroom"
	} else {
		reply["cmd"] = "unlockroom"
	}
	c.SendJSON(reply)
}
