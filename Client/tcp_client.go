package Client

import (
	"appServer/Config"
	"appServer/Error"
	"io/ioutil"
	"log"
	"net"
)

type Client struct {
	client_host     string
	client_port     string
	client_protocol string
}

func Create_client() *Client {
	log.Println("Creating client:")
	return &Client{}
}

func (c *Client) Configure_client(server_config Config.MyServer) {

	log.Println("Getting configuration:")
	host, port, protocol := server_config.GetConfig()

	log.Println("Configuring the client:")
	c.client_host = host
	c.client_port = port
	c.client_protocol = protocol
}

func (c *Client) Get_connection() net.Conn {

	log.Println("Creating connection object:")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", c.client_host+":"+c.client_port)
	Error.Handle_error(err)
	conn, err := net.DialTCP(c.client_protocol, nil, tcpAddr)
	Error.Handle_error(err)
	return conn
}

func (c *Client) Transmitt_data(message string) string {

	log.Println("Creating connection with the remote server:")
	conn := c.Get_connection()
	num, err := conn.Write([]byte(message))

	log.Println("Number of bytes read:", num)
	Error.Handle_error(err)

	log.Println()
	result, err := ioutil.ReadAll(conn)
	Error.Handle_error(err)
	return string(result)
}
