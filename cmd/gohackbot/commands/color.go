package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0x17de/gohackchat/pkg/hack"
)

type ColorCommand struct {
	Command
}

func NewColorCommand() *ColorCommand {
	return &ColorCommand{Command{
		aliases:     []string{"color"},
		description: "Set a user color",
	}}
}

func (cmd *ColorCommand) Run(c *hack.Client, root hack.JsonValue) {
	args := cmd.GetArgs(root)
	reply := make(hack.JsonValue)
	if len(args) < 2 {
		return
	}

	requestedColor := strings.ToLower(args[1])
	reply["cmd"] = "forcecolor"
	switch len(args) {
	case 2:
		trip, tripok := root["trip"].(string)
		if !tripok {
			return
		}

		permissionsFile, err := os.Open("colorperms")
		if err != nil {
			return
		}

		var colors []string
		scanner := bufio.NewScanner(permissionsFile)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, trip) {
				colors = strings.Split(line, " ")[1:]
			}
		}

		allowedColor := false
		for _, color := range colors {
			if requestedColor == color {
				allowedColor = true
				break
			}
		}

		if !allowedColor {
			return
		}
		fmt.Printf("Setting %s#%s's color to %s", root["nick"], trip, requestedColor)

		reply["nick"] = root["nick"]
		reply["color"] = requestedColor
	case 3:
		if !c.IsMod(root) {
			return
		}
		reply["nick"] = requestedColor
		reply["color"] = args[2]
	}
	c.SendJSON(reply)
}
