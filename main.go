package main

import (
	"appServer/Config"
	"appServer/ServerLogic"
	"appServer/SocketUtil"
	"fmt"
	"log"
)

func server_test() {

	log.Println("Implementing time server:")
	config_obj := Config.MyServer{}

	config_obj.AddConfig("localhost", "9000", "tcp")
	fmt.Println(config_obj.GetConfig())

	log.Println("Creating the server:")
	server := SocketUtil.CreateServer()

	log.Println("Configuring the time server:")
	server.ConfigureServer(config_obj)

	log.Println("Running the time server:")
	server.Run_server_logic(ServerLogic.TimeServer)
}
