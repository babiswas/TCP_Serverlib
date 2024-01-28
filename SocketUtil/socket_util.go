package SocketUtil

import (
	"appServer/Config"
	"appServer/Error"
	"log"
	"net"
)

type Server struct {
	server_host     string
	server_port     string
	server_protocol string
}

func CreateServer() *Server {

	log.Println("Creating a server:")
	return &Server{}
}

func (s *Server) ConfigureServer(server_config Config.MyServer) {

	log.Println("Fetching data from configurations:")

	host, port, protocol := server_config.GetConfig()
	log.Println("HOST: ", host, " PORT: ", port, " PROTOCOL: ", protocol)

	s.server_host = host
	s.server_port = port
	s.server_protocol = protocol
}

func (s *Server) Create_Listener() net.Listener {

	log.Println("Creating a listener:")
	listen, err := net.Listen(s.server_protocol, s.server_host+":"+s.server_port)
	Error.Handle_error(err)

	log.Println("Sucessfully created a listener.")
	return listen
}

func (s *Server) Run_server_logic(server_logic func(conn net.Conn)) {

	log.Println("Creating listener for running server logic:")
	listener := s.Create_Listener()

	defer listener.Close()

	log.Println("Running infinite loop for receiving connections:")
	for {
		conn, err := listener.Accept()
		Error.Handle_error(err)
		go server_logic(conn)
	}
}
