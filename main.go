package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/deepam02/go-file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt the file")
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("file encyption and decryption tool")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\t go run . encrypt /path/to/file")
	fmt.Println("\t go run . decrypt /path/to/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\t encrypt\tencrypts a file")
	fmt.Println("\t decrypt\tdecrypts a file")
	fmt.Println("\t help\t\tprints help")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to encrypt")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}
	password := getPassword()
	fmt.Println("\n encrypting ...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file encrypted successfully")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to decrypt")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}
	fmt.Print("enter password:")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("\n decrypting ...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file decrypted successfully")
}

func getPassword() []byte {
	fmt.Print("enter password:")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Print("\n confirm password:")
	confirmPassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	if !validatePassword(password, confirmPassword) {
		fmt.Println("passwords do not match")
		return getPassword()
	}
	return password

}

func validatePassword(password []byte, confirmPassword []byte) bool {
	return bytes.Equal(password, confirmPassword)

}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
