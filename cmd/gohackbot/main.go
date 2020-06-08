package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	cmds "github.com/0x17de/gohackchat/cmd/gohackbot/commands"
	hack "github.com/0x17de/gohackchat/pkg/hack"
)

func getLogin(username, password string) (login string) {
	if password != "" {
		login = fmt.Sprintf("%s#%s", username, password)
	} else {
		login = username
	}
	return
}

func main() {
	username := flag.String("username", "gohackbot", "Username")
	password := flag.String("password", "", "Password")
	passwordFile := flag.String("password-file", "", "Read the password from file")
	channel := flag.String("channel", "botDev", "Password")
	prefix := flag.String("prefix", "&", "The command prefix")
	flag.Parse()

	if len(*prefix) != 1 {
		log.Fatalf("Only one character prefixes are supported yet")
	}

	if *password == "" && *passwordFile != "" {
		data, err := ioutil.ReadFile(*passwordFile)
		*password = string(data)
		if err != nil {
			log.Fatalf("Failed to read password file: %v", err)
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	client, err := hack.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	helpCommand := cmds.NewHelpCommand()

	userCommandModule := cmds.NewCommandModule(*prefix, false)
	userCommandModule.Register(cmds.NewTestCommand())
	userCommandModule.Register(cmds.NewColorCommand())
	userCommandModule.Register(helpCommand)

	modCommandModule := cmds.NewCommandModule(*prefix, true)
	modCommandModule.Register(cmds.NewLockRoomCommand())
	modCommandModule.Register(cmds.NewCaptchaCommand())
	modCommandModule.Register(cmds.NewMuteCommand())
	modCommandModule.Register(cmds.NewKickCommand())
	modCommandModule.Register(cmds.NewAuthCommand())

	helpCommand.Register("Mod", modCommandModule)
	helpCommand.Register("User", userCommandModule)

	client.Register(NewPrintModule())
	client.Register(modCommandModule)
	client.Register(userCommandModule)

	client.JoinChannel(*channel, getLogin(*username, *password))
	go client.Run()

	for {
		select {
		case <-client.C:
			return // terminate
		case <-interrupt:
			client.Stop()
			break
		}
	}
}
