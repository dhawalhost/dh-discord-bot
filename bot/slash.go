package bot

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @file Description
 * @desc Created on 2021-05-02 10:16:25 pm
 * @copyright Crearosoft
 */

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// GuildID - server ID
	GuildID  string
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "gettime",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Get server date and time",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"gettime": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			dt := time.Now()
			fdt := fmt.Sprintf("Current date and time is: %s", dt.String())
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: fdt,
				},
			})
		},
	}
)

// InitiateCommonHandler -
func InitiateCommonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.Data.Name]; ok {
		h(s, i)
	}
	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
