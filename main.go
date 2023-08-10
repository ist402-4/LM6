package main

import (
	"fmt"
	"strings"
)

var rotors = []string{
	"JGDQOXUSCAMIFRVTPNEWKBLZYH",
	"NTZPSFBOKMWRCJDIVLAEYUXHGQ",
	"JVIUBHTCDYAKEQZPOSGXNRMWFL",
}
var reflector = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var plugboard = map[rune]rune{
	'A': 'C', 'B': 'F', 'C': 'A', 'D': 'E', 'E': 'D', 'F': 'B', 'G': 'G',
}

func main() {
	message := "Hello"
	encrypted := processMessage(message, true)
	decrypted := processMessage(encrypted, false)
	fmt.Println("Original Message:", message)
	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)
}
func processMessage(text string, isEncrypting bool) string {
	text = strings.ToUpper(text)
	var processedText strings.Builder
	//https://golangdocs.com/blank-identifier-in-golang - BLANK IDENTIFIER
	for _, char := range text {
		if plug, exists := plugboard[char]; exists {
			char = plug
		}
		rotorFunc := performSubstitution
		if !isEncrypting {
			rotorFunc = performReverseSubstitution
		}
		for _, rotor := range rotors {
			char = rotorFunc(char, rotor)
		}
		if isEncrypting {
			char = performSubstitution(char, reflector)
		} else {
			char = performReverseSubstitution(char, reflector)
		}
		for i := len(rotors) - 1; i >= 0; i-- {
			char = rotorFunc(char, rotors[i])
		}
		if plug, exists := plugboard[char]; exists {
			char = plug
		}
		processedText.WriteRune(char)
	}
	return processedText.String()
}
func performSubstitution(char rune, rotor string) rune {
	return rune(rotor[char-'A'])
}
func performReverseSubstitution(char rune, rotor string) rune {
	index := strings.IndexRune(rotor, char)
	return rune(index + 'A')
}
