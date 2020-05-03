package commands

import (
	"github.com/0x17de/gohackchat/pkg/hack"
)

type AuthCommand struct {
	Command
}

func NewAuthCommand() *AuthCommand {
	return &AuthCommand{Command{
		aliases:     []string{"authtrip", "deauthtrip"},
		description: "Bypass channel locks/capcha when trip matches",
	}}
}

func (cmd *AuthCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)
	if len(args) < 2 {
		return
	}

	reply := make(hack.JsonValue)
	if args[0] == "authtrip" {
		reply["cmd"] = "authtrip"
	} else {
		reply["cmd"] = "deauthtrip"
	}
	reply["trip"] = args[1]
	c.SendJSON(reply)
}
