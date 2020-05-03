package commands

import (
	"github.com/0x17de/gohackchat/pkg/hack"
)

type KickCommand struct {
	Command
}

func NewKickCommand() *KickCommand {
	return &KickCommand{Command{
		aliases:     []string{"kick"},
		description: "Kick a user from the channel",
	}}
}

func (cmd *KickCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)
	if len(args) < 2 {
		return
	}

	reply := make(hack.JsonValue)
	reply["cmd"] = "kick"
	reply["nick"] = args[1]
	c.SendJSON(reply)
}
