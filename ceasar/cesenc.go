package main

import "fmt"

func cipher(input string, shift int) string {
	encrypted := []rune(input)

	for i, char := range encrypted {
		if char >= 'a' && char <= 'z' {
			encrypted[i] = 'a' + (char-'a'+rune(shift))%26
		} else if char >= 'A' && char <= 'Z' {
			encrypted[i] = 'A' + (char-'A'+rune(shift))%26
		}

	}
	return string(encrypted)
}

func main() {
	fmt.Print("Enter text to encrypt: ")

	var message string
	fmt.Scanln(&message)
	fmt.Print("Enter shift value: ")

	var shift int
	fmt.Scanln(&shift)

	encryptedMessage := cipher(message, shift)
	fmt.Println("Encrypted message:", encryptedMessage)
}
