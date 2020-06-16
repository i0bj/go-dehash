package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"os"
)

var (
	md5hash    string
	sha256hash string
)

func MD5hasher() {
	scanMd5 := bufio.NewScanner(wl) // TODO handle wl file
	for scanMd5.Scan() {
		password := scanMd5.Text()
		// Hashing password read from password file.
		Mhash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if md5hash == Mhash {
			fmt.Printf("[!] Password Found!\nmd5: %s\n", password)
		}
	}
}

func SHA256hasher() {
	scanSHA256 := bufio.NewScanner(wl) // TODO handle wl file
	for scanSHA256.Scan() {
		password := scanSHA256.Text()
		S256hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if sha256hash == S256hash {
			fmt.Printf("[!] Password Found!\nSHA256:\n", password)
		}
	}

}

func main() {
	// Scanning for user input
	fmt.Println("1. MD5")
	fmt.Println("2. SHA256")
	// Password list I found online. Will be used to Match digest.
	wl, err := os.Open("passwordlist.txt")
	if err != nil {
		fmt.Println("Error when opening passwordlist.txt")
	}
	defer wl.Close() // Clean up, close file to save resources.

	var selection int
	switch selection {
	case 1:
		MD5hasher()
	case 2:
		SHA256hasher()
	case 3:
		fmt.Println("Exiting...")
		os.Exit(2)
	}

}
