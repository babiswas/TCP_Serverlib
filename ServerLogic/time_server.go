package ServerLogic

import (
	"appServer/Error"
	"fmt"
	"log"
	"net"
	"time"
)

func TimeServer(conn net.Conn) {

	defer conn.Close()

	log.Println("Creating buffer:")
	buffer := make([]byte, 1024)
	num, err := conn.Read(buffer)

	Error.Handle_error(err)
	fmt.Println("Number of bytes read:", num)

	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")

	log.Println("Current Time:", time)

	conn.Write([]byte("Hi back!"))
	conn.Write([]byte(time))
}
