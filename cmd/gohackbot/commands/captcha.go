package commands

import (
	"github.com/0x17de/gohackchat/pkg/hack"
)

type CaptchaCommand struct {
	Command
}

func NewCaptchaCommand() *CaptchaCommand {
	return &CaptchaCommand{Command{
		aliases:     []string{"enablecaptcha", "disablecaptcha"},
		description: "Enable/disable captcha in current channel",
	}}
}

func (cmd *CaptchaCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)

	reply := make(hack.JsonValue)
	if args[0] == "enablecaptcha" {
		reply["cmd"] = "enablecaptcha"
	} else {
		reply["cmd"] = "disablecaptcha"
	}
	c.SendJSON(reply)
}
