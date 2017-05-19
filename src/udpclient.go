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

// Error check
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	//kobler opp til serveren
	conn, err := net.Dial("udp", HOST+":"+PORT)
	CheckError(err)

	defer conn.Close()

	for {
		//Ny reader som lar deg skrive melding fra terminalen
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Skriv til server: ")
		tekst, _ := reader.ReadString('\n')

		//krypterer teksten
		kryptert := krypter(tekst)

		//sender meldingen
		_, err = conn.Write(kryptert)
		CheckError(err)
		fmt.Println("")
		fmt.Println("Melding kryptert og sendt")
		melding := string(kryptert)

		//For å se at meldingen er blitt kryptert
		fmt.Println("Melding etter kryptering: " + melding)
		fmt.Println("")


	}
}














