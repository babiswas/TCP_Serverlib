package main

import (
	"appServer/Client"
	"appServer/Config"
	"fmt"
	"log"
)

func main() {

	log.Println("Create server configuration:")
	server_config := Config.MyServer{}
	server_config.AddConfig("localhost", "9000", "tcp")

	log.Println("Creating client:")
	client := Client.Create_client()
	client.Configure_client(server_config)

	log.Println("Connecting to remote server:")
	fmt.Println(client.Transmitt_data("hello"))
}
