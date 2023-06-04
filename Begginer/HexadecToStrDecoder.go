// Decoding Messages

// In a cyberpunk world, secret messages are encoded in a unique format and it's your job to decode them. Each message is a slice
// of strings, and each string contains only hexadecimal values. Each hex value, when converted to ASCII, corresponds to a
//character in the secret message.

// Write a Go program to convert each hexadecimal value into ASCII, then concatenate the ASCII characters to get the full message.

// Input: []string{"48", "65", "6c", "6c", "6f", "20", "57", "6f", "72", "6c", "64"}
// Output: Hello World

package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	secretMessage := []string{"48", "65", "6c", "6c", "6f", "20", "57", "6f", "72", "6c", "64"}
	decodedSecretMessage := Hex2StrDecoder(secretMessage)
	fmt.Printf("The decoded secret message is: %v\n", decodedSecretMessage)
}

func Hex2StrDecoder(arr []string) string {
	var message string
	for _, hexval := range arr {
		byteVal, err := hex.DecodeString(hexval)
		if err != nil {
			log.Fatal(err)
		}
		message += string(byteVal)
	}
	return message
}
