package main

import (
	"log"
	"strings"

	hack "github.com/0x17de/gohackchat/pkg/hack"
)

type PrintModule struct {
}

func NewPrintModule() *PrintModule {
	return &PrintModule{}
}
func (m *PrintModule) OnMessage(c *hack.Client, root hack.JsonValue) bool {
	nick, ok := root["nick"].(string)
	if !ok {
		log.Printf("Chat message is missing field: nick")
		return true
	}
	text, ok := root["text"].(string)
	if !ok {
		log.Printf("Chat message is missing field: text")
		return true
	}
	trip, ok := root["trip"].(string)
	if ok {
		trip = "#" + trip
	} else {
		trip = ""
	}
	mod := c.IsMod(root)

	var userType string
	if mod {
		userType = "M"
	} else {
		userType = "U"
	}

	log.Printf("%s %s%s: %s", userType, nick, trip, text)
	return true
}

func (m *PrintModule) OnWhisper(_ *hack.Client, root hack.JsonValue) bool {
	nick, ok := root["from"].(string)
	if !ok {
		log.Printf("Whisper is missing field: from")
		return true
	}
	text, ok := root["text"].(string)
	if !ok {
		log.Printf("Whisper is missing field: text")
		return true
	}
	substrings := strings.SplitN(text, ": ", 2)
	if len(substrings) < 2 {
		log.Printf("Whisper text is not divided by :")
		return true
	}
	log.Printf("%s whispered: %s", nick, substrings[1])
	return true
}
