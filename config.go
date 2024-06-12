// this file only revolve aroound "config.json" file  to read it

package config

import (
	// we use "enconfig/json" because we want to use the unmarshalling function from json package
	"encoding/json"

	"fmt"

	// we use io/ioutil because it help to read the file

	"io/ioutil"
)

// here we define  3 variables that will be use in code

var (
	Token     string
	BotPrefix string

	config *configStruct
)

// configStruct struct are  datatype that create an own datatype

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {

	fmt.Println("Reading config file...")

	// here we use the ioutil function to read the config file

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))

	// now we do the unmarshel the file
	err = json.Unmarshal(file, config)
	if err == nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}
