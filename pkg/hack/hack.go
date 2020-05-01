package hack

import (
	"encoding/json"
	"log"
	"time"

	ws "github.com/gorilla/websocket"
)

type JsonValue = map[string]interface{}

type Client struct {
	cxn     *ws.Conn
	modules []interface{}
	done    chan int
	C       chan struct{}
}

type MessageModule interface {
	OnMessage(client *Client, root JsonValue) bool
}
type WhisperModule interface {
	OnWhisper(client *Client, root JsonValue) bool
}

func NewClient() (*Client, error) {
	cxn, _, err := ws.DefaultDialer.Dial("wss://hack.chat/chat-ws", nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		cxn:     cxn,
		modules: make([]interface{}, 0),
		done:    make(chan int),
		C:       make(chan struct{})}, nil
}

func (c *Client) Register(m interface{}) {
	c.modules = append(c.modules, m)
}

func (c *Client) Run() {
	defer close(c.C)

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	go func() {
		defer c.Stop()
		c.runMessageLoop()
	}()

	for {
		select {
		case <-c.done:
			return
		case <-ticker.C:
			c.onPing()
		}
	}
}

func (c *Client) Stop() {
	c.done <- 1
}

func (c *Client) SendJSON(json JsonValue) {
	c.cxn.WriteJSON(json)
}

func (c *Client) SendMessage(message string) {
	root := make(JsonValue)
	root["cmd"] = "chat"
	root["text"] = message
	c.SendJSON(root)
}

func (c *Client) IsMod(json JsonValue) bool {
	mod, ok := json["mod"].(bool)
	if !ok {
		return false
	}
	return mod
}

func (c *Client) runMessageLoop() {
	for {
		_, message, err := c.cxn.ReadMessage()
		if err != nil {
			log.Println("read: %v", err)
			return
		}
		c.onMessage(message)
	}
}

func (c *Client) onPing() {
	root := make(JsonValue)
	root["cmd"] = "ping"
	c.SendJSON(root)
}

func (c *Client) onMessage(message []byte) {
	var root interface{}
	err := json.Unmarshal(message, &root)
	if err != nil {
		log.Printf("Failed to parse json: %s", message)
	}

	switch root.(type) {
	case JsonValue:
		rootobj := root.(JsonValue)
		switch rootobj["cmd"] {
		case "onlineSet":
			c.onUserList(rootobj)
			break
		case "chat":
			c.onChatMessage(rootobj)
			break
		case "info":
			if rootobj["type"] == "whisper" {
				c.onChatWhisper(rootobj)
			}
			break
		default:
			log.Printf("Unhandled: %s", message)
			break
		}
		break
	default:
		log.Printf("recv: %s", message)
		log.Printf("The root element is not an object: %T", root)
		break
	}
}

func (c *Client) onChatMessage(root JsonValue) {
	for _, module := range c.modules {
		if m, ok := module.(MessageModule); ok {
			if !m.OnMessage(c, root) {
				break
			}
		}
	}
}

func (c *Client) onChatWhisper(root JsonValue) {
	for _, module := range c.modules {
		if m, ok := module.(WhisperModule); ok {
			if !m.OnWhisper(c, root) {
				break
			}
		}
	}
}

func (c *Client) onUserList(root JsonValue) {
}

func (c *Client) JoinChannel(channel string, nick string) {
	root := make(JsonValue)
	root["cmd"] = "join"
	root["channel"] = channel
	root["nick"] = nick
	c.SendJSON(root)
}
