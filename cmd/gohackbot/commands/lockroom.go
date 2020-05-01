package commands

import (
	"github.com/0x17de/gohackchat/pkg/hack"
)

type LockRoomCommand struct {
	lock bool
}

func NewLockRoomCommand(lock bool) *LockRoomCommand {
	return &LockRoomCommand{lock: lock}
}

func (cmd *LockRoomCommand) Run(c *hack.Client, root hack.JsonValue) {
	reply := make(hack.JsonValue)
	if cmd.lock {
		reply["cmd"] = "lockroom"
	} else {
		reply["cmd"] = "unlockroom"
	}
	c.SendJSON(reply)
}
