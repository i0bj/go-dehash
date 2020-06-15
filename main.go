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

	/* Opening word list I found online
	   this word list will each password will
	   be hashed and checked against the provided hash
	   if match found password found.*/
	wl, err := os.Open("passwordlist.txt")
	if err != nil {
		fmt.Println("Error when opening passwordlist.txt")
	}
	defer wl.Close() // Clean up, close file to save resources.

	scanner := bufio.NewScanner(wl)
	for scanner.Scan() {
		password := scanner.Text()
		// Hashing password read from password file.
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == md5hash {
			fmt.Printf("Password found!\n md5:\n%s", password)
		} else {
			fmt.Println("Password not found...")
		}
	}
}
