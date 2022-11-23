package main

import (
	"os"

	"github.com/Drambluker/godirect/config"
	"github.com/Drambluker/godirect/server"
)

func main() {
	configArg := os.Args[1:][0]
	config := config.NewConfig(configArg)
	server := server.NewServer(*config)
	server.Run()
}
