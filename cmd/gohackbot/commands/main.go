package commands

import (
	"strings"

	hack "github.com/0x17de/gohackchat/pkg/hack"
)

type ICommand interface {
	GetArgs(root hack.JsonValue) []string
	GetAliases() []string
	GetDescription() string
	Run(client *hack.Client, root hack.JsonValue)
}

type Command struct {
	aliases     []string
	description string
}

func (c *Command) GetAliases() []string {
	return c.aliases
}

func (c *Command) GetDescription() string {
	return c.description
}

func (c *Command) GetArgs(root hack.JsonValue) []string {
	text, ok := root["text"].(string)
	if !ok {
		return nil
	}

	return strings.Split(text[1:], " ")
}

type CommandModule struct {
	prefix    string
	commands  map[string]ICommand
	mustBeMod bool
}

func NewCommandModule(prefix string, mustBeMod bool) *CommandModule {
	return &CommandModule{
		prefix:    prefix,
		commands:  make(map[string]ICommand),
		mustBeMod: mustBeMod,
	}
}

func (m *CommandModule) Register(cmd ICommand) {
	for _, name := range cmd.GetAliases() {
		m.commands[name] = cmd
	}
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
