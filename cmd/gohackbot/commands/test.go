package commands

import "github.com/0x17de/gohackchat/pkg/hack"

type TestCommand struct {
	Command
}

func NewTestCommand() *TestCommand {
	return &TestCommand{Command{
		aliases:     []string{"test"},
		description: "A simple reply test",
	}}
}

func (t *TestCommand) Run(c *hack.Client, root hack.JsonValue) {
	c.SendMessage("Works!")
}
