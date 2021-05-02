package bot

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @file Description
 * @desc Created on 2021-05-02 3:26:18 pm
 * @copyright Crearosoft
 */
import (
	"fmt"

	"github.com/dhawalhost/dh-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

// BotID is a unique bot ID
var BotID string
var goBot *discordgo.Session

// Start function makes the bot online
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		res, _ := s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
