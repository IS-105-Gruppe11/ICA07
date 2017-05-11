package main

import (
	"fmt"
	"net"
	"os"
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

	ServerAddr,err := net.ResolveUDPAddr("udp",":"+PORT)
	CheckError(err)
	// listen to incoming udp packets
	ServerConn, err := net.ListenUDP("udp", ServerAddr )
	CheckError(err)
	defer ServerConn.Close()

	// reader for serveren
	buffer := make([]byte, 1024)

	for {
		n,addr,err := ServerConn.ReadFromUDP(buffer)

		// printer ut melding fra clienten
		fmt.Println("Recieved ", string(buffer[0:n]), " from ", addr)
		CheckError(err)
	}
}
