package commands

import (
	"regexp"

	"github.com/0x17de/gohackchat/pkg/hack"
)

type MuteCommand struct {
	Command
	ipRE *regexp.Regexp
}

func NewMuteCommand() *MuteCommand {
	return &MuteCommand{
		Command: Command{
			aliases:     []string{"mute", "unmute"},
			description: "(Un-)Mute a user",
		},
		ipRE: regexp.MustCompile(`^[0-9]+(\.[0-9]+){3}$`),
	}
}

func (cmd *MuteCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)
	if len(args) < 2 {
		return
	}

	reply := make(hack.JsonValue)
	if args[0] == "mute" {
		reply["cmd"] = "dumb"
		if len(args) > 2 {
			reply["allies"] = args[2:]
		}
		reply["nick"] = args[1]
	} else {
		reply["cmd"] = "speak"

		if cmd.ipRE.MatchString(args[1]) {
			reply["ip"] = args[1]
		} else {
			reply["hash"] = args[1]
		}
	}
	c.SendJSON(reply)
}
