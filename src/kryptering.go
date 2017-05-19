package main

import (
	"crypto/rand"
	"io"
	"golang.org/x/crypto/nacl/secretbox"
	"encoding/hex"
)

const (
	PASSWORD = "3571466f198420646869732000117373766e3264207461296120736563726575"
)
// funksjon for kryptering
func krypter(message string) []byte {

	secretKeyBytes, err := hex.DecodeString(PASSWORD)
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	// krypterer meldingen
	encrypted := secretbox.Seal(nonce[:], []byte(message), &nonce, &secretKey)
	return encrypted
}
//funksjon for dekryptering
func dekrypter(encrypted []byte) string  {

	secretKeyBytes, err := hex.DecodeString(PASSWORD)
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error")
	}
	return string(decrypted)
}
