package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

const (
	HOST = "127.0.0.1"
	PORT = "8001"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func main() {

	ServerAddr,err := net.ResolveTCPAddr("tcp",":"+PORT)
	CheckError(err)
	// listen to incoming udp packets
	ServerConn, err := net.ListenTCP("tcp", ServerAddr )
	CheckError(err)
	defer ServerConn.Close()
	CheckError(err)

	for {
		//accept connections using Listener.Accept()
		c, err := ServerConn.Accept()
		CheckError(err)

		// goroutine behandler tilkoblingen
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {

	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		// output message received
		fmt.Print("Melding:", string(message))

	}

}
	


