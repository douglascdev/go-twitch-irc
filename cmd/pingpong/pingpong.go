package main

import (
	"log"
	"strings"

	//nolint:depguard
	twitch "github.com/gempir/go-twitch-irc/v4"
)

const (
	clientUsername            = "justinfan123123"
	clientAuthenticationToken = "oauth:123123123"
)

func main() {
	client := twitch.NewClient(clientUsername, clientAuthenticationToken)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.Contains(strings.ToLower(message.Message), "ping") {
			log.Println(message.User.Name, "PONG", message.Message)
		}
	})

	client.Join("testaccount_420")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
