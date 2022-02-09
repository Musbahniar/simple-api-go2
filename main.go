package main

import (
	"fmt"
	"log"
	"os"
	"simple-api-go2/api/utils"
	"simple-api-go2/config"

	"github.com/err-him/gonf"
)

type PortConfig struct {
	Port string
}

func main() {

	//Get port from command line interface
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}
	cfg := PortConfig{}
	err := gonf.GetConfig(utils.GetEnvFile(env), &cfg)
	if err != nil {
		log.Fatal("environment can not be loaded at this moment, Please try after some time", err)
	}
	app := &config.App{}
	fmt.Println("App run on port: " + cfg.Port)
	app.Intialize()
	app.Run(cfg.Port)
}
