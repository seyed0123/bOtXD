package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("src/config.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(file))

	config = &configStruct{}
	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
