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
	conn, err := net.Dial("udp", HOST+":"+PORT)
	CheckError(err)

	defer conn.Close()

	for {

		reader := bufio.NewReader(os.Stdin)   //lager ny reader
		fmt.Print("Skriv til server: ")
		text, _ := reader.ReadString('\n')
		message := []byte(text)

		_, err = conn.Write(message)    //sender meldingen
		CheckError(err)
	}
}














