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
		tekst, _ := reader.ReadString('\n')

		//krypterer teksten
		kryptert := krypter(tekst)
		_, err = conn.Write(kryptert)    //sender meldingen
		CheckError(err)
		fmt.Println("Melding kryptert og sendt")


	}
}














