package main

import (
	"strings"

	hack "github.com/0x17de/gohackchat/pkg/hack"
)

type Command interface {
	Run(client *hack.Client, root hack.JsonValue)
}

type CommandModule struct {
	prefix    string
	commands  map[string]Command
	mustBeMod bool
}

func NewCommandModule(prefix string, mustBeMod bool) *CommandModule {
	return &CommandModule{
		prefix:    prefix,
		commands:  make(map[string]Command),
		mustBeMod: mustBeMod,
	}
}

func (m *CommandModule) Register(name string, cmd Command) {
	m.commands[name] = cmd
}

func (m *CommandModule) OnMessage(client *hack.Client, root hack.JsonValue) bool {
	var text string
	text, ok := root["text"].(string)
	if !ok {
		return true
	}
	if m.mustBeMod && !client.IsMod(root) {
		return true
	}
	if strings.HasPrefix(text, m.prefix) {
		for name, cmd := range m.commands {
			if strings.HasPrefix(text[1:], name) && (len(text) == len(name)+1 || text[len(name)+1] == ' ') {
				cmd.Run(client, root)
				return false
			}
		}
	}
	return true
}

type TestCommand struct {
}

func (t *TestCommand) Run(c *hack.Client, root hack.JsonValue) {
	c.SendMessage("Works!")
}
