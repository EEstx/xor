package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to cipher or decipher, def is cipher")
var secretKey = flag.String("secret", "", "Your secret key")

func main() {
	flag.Parse()

	if len(*secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "No secret is provided! Exiting now...")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter your text to cipher: ")
		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to encrypt %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cipheredText)
	case "decipher":
		cipherdtext := getUserInput("Enter your text to decipher: ")
		decipheredText, err := cipherer.Decipher(cipherdtext, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to decrypt %v\n", err)
			os.Exit(1)
		}
		fmt.Println(decipheredText)
	default:
		fmt.Println("Invalid mode")
		os.Exit(1)
	}
	if *secretKey == "" {
		fmt.Println("Please enter a secret key")
		os.Exit(1)
	}

}
func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something wrong")
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}
