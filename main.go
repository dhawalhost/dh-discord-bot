package main

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @file Description
 * @desc Created on 2021-05-02 3:20:32 pm
 * @copyright Crearosoft
 */

import (
	"fmt"

	"github.com/dhawalhost/dh-discord-bot/bot"
	"github.com/dhawalhost/dh-discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	bot.Start()
	<-make(chan struct{})
	return
}
