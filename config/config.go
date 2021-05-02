package config

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @file Description
 * @desc Created on 2021-05-02 4:01:47 pm
 * @copyright Crearosoft
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botprefix"`
}

func ReadConfig() error {
	fmt.Println("Reading from config file ...")
	file, err := ioutil.ReadFile("./config/config.json")

	if err != nil {
		// fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)

	if err != nil {
		// fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
