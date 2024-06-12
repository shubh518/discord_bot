// this file is responsible for sending messags to discord
package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"main.go/config"
)

var BotId string

//here we create a variable called goBot which will be type session

var goBot *discordgo.Session

func Start() {
	// here  we create a new session "discordgo.new" and pass the "config.Token"
	goBot, err := discordgo.New("Bot" + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	// here we have to create a bot as a user

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	// since we have a user we want to store a  BotId

	BotId = u.ID

	// the other function from the goBot library  that we want to use  is the :-
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running...")

}

// here we create messageHandler function
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {

		return

	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
