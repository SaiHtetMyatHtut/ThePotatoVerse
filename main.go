package main

import (
	"log"

	"github.com/SaiHtetMyatHtut/potatoverse/configs"
	"github.com/SaiHtetMyatHtut/potatoverse/infra/di"
	"github.com/SaiHtetMyatHtut/potatoverse/src/server"
)

func main() {
	configs.LoadEnv()

	container := di.Initialize()
	err := container.Invoke(server.NewServer)
	if err != nil {
		log.Fatal(err)
	}
}
