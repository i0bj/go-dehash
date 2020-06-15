package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	// Scanning for user input
	fmt.Println("Enter the MD5 hash: ")
	var md5hash string
	fmt.Scanln(&md5hash)

	wl, err := os.Open("passwordlist.txt")
	if err != nil {
		fmt.Println("Error when opening passwordlist.txt")
	}
	defer wl.Close()

	scanner := bufio.NewScanner(wl)
	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == md5hash {
			fmt.Printf("Password found!\n md5:\n%s", password)
		}
	}
}
