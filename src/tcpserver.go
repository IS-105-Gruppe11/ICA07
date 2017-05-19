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
	//Forbereder tilkobling til port ardressen
	ServerAddr,err := net.ResolveTCPAddr("tcp",":"+PORT)
	CheckError(err)

	// Hører etter tcp pakker
	ServerConn, err := net.ListenTCP("tcp", ServerAddr )
	CheckError(err)
	defer ServerConn.Close()
	CheckError(err)

	for {
		//Akstepterer tilkoblingen fra klient
		c, err := ServerConn.Accept()
		CheckError(err)

		// goroutine behandler tilkoblingen
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {

	for {

		melding, _ := bufio.NewReader(c).ReadString('\n')
		// printer ut melding fra klient
		fmt.Print("Melding:", melding)
	}



}
	


