package Config

import "log"

type MyServer struct {
	host     string
	port     string
	protocol string
}

func (m *MyServer) AddConfig(host, port, protocol string) {

	log.Println("Creating a configurations:")

	m.host = host
	m.port = port
	m.protocol = protocol

	log.Println("Config: HOST->", host, "PORT->", port, "PROTOCOL->", protocol)
}

func (m MyServer) GetConfig() (string, string, string) {
	return m.host, m.port, m.protocol
}
