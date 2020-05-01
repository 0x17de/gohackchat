package commands

import (
	"strings"

	"github.com/0x17de/gohackchat/pkg/hack"
)

type HelpCommand struct {
	Command
	commandModules map[string]*CommandModule
}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{
		Command: Command{
			aliases:     []string{"help"},
			description: "List available commands",
		},
		commandModules: make(map[string]*CommandModule),
	}
}

func (cmd *HelpCommand) Register(name string, module *CommandModule) {
	cmd.commandModules[name] = module
}

func (cmd *HelpCommand) Run(c *hack.Client, root hack.JsonValue) {
	var sb strings.Builder
	for name, module := range cmd.commandModules {
		sb.WriteString(name)
		sb.WriteString(" commands:")
		for _, command := range module.commands {
			sb.WriteString("\n  ")
			sb.WriteString(strings.Join(command.GetAliases(), ", "))
			sb.WriteString(": ")
			sb.WriteString(command.GetDescription())
		}
		sb.WriteString("\n")
	}
	c.SendMessage(sb.String())
}
