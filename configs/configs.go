package configs

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	/*
		Struct to holds the env values.
	*/
	env struct {
		Redis  redis  `json:"redis"`
		Server server `json:"server"`
	}

	redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	server struct {
		Port string `json:"port"`
	}
)

var Env env

/*
Loads env configs from the config file. If there is an error while loading the configs,
an error will be rasied.
*/
func LoadEnv() {
	_, err := toml.DecodeFile("./configs/configs.toml", &Env)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
