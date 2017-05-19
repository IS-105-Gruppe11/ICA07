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


func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	//Kobler opp til serveren
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	CheckError(err)

	defer conn.Close()

	for {
		//lager ny reader
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Skriv melding her: ")
		text, _ := reader.ReadString('\n')

		//Gjør meldingen om til byteslice
		message := []byte(text)

		//sender meldingen
		_, err = conn.Write(message)
		CheckError(err)
	}
}

