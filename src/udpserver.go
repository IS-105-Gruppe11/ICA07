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
	// Forbereder tilkoblingen til port adressen
	ServerAddr,err := net.ResolveUDPAddr("udp",":"+PORT)
	CheckError(err)
	// Hører etter udp pakker
	ServerConn, err := net.ListenUDP("udp", ServerAddr )
	CheckError(err)
	defer ServerConn.Close()

	// reader for serveren
	buffer := make([]byte, 1024)

	fmt.Println("Server igang og klar til bruk...")

	//En loop som tar imot udp pakker
	for {
		n,addr,err := ServerConn.ReadFromUDP(buffer)

		//konverterer byteslicene til string
		melding := buffer[0:n]

		//dekrypterer meldingen
		dekryptert := dekrypter(melding)

		// printer ut melding fra klienten
		fmt.Println("Motatt melding: ", dekryptert, " fra ", addr)
		CheckError(err)
	}
}
