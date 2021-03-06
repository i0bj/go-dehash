package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// goroutine to capture input from user to exit program.
// can add additional code in function to clean up.
func programExit() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r[!] Ctrl+C pressed. Program exiting..")
		os.Exit(0)
	}()
}

func main() {

	var md5hash string
	var sha256hash string

	programExit()

	wl, err := os.Open("passwordlist.txt")
	if err != nil {
		fmt.Println("Error when opening passwordlist.txt")
	}

	defer wl.Close() // Clean up, close file to save resources.
	// Scanning for user input
	fmt.Println("[+] Select the hash whould you like to match: \n ")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA256")
	// Password list I found online. Will be used to Match digest.

	var selection int
	for ok := true; ok; ok = (selection != 2) {
		n, err := fmt.Scanln(&selection)
		if n > 1 || err != nil {
			fmt.Println("[!] Invalid Selection")
			continue
		}
		switch selection {
		case 1:
			fmt.Println("Enter md5 Hash")
			fmt.Scanln(&md5hash)
			scanMd5 := bufio.NewScanner(wl)
			for scanMd5.Scan() {
				password := scanMd5.Text()
				// Hashing password read from password file.
				Mhash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
				if md5hash == Mhash {
					fmt.Printf("[!] Password Found!\nmd5: %s\n", password)
				}
			}
		case 2:
			fmt.Println("Enter SHA256 Hash")
			fmt.Scanln(&sha256hash)
			scanSHA256 := bufio.NewScanner(wl)
			for scanSHA256.Scan() {
				password := scanSHA256.Text()
				S256hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
				if sha256hash == S256hash {
					fmt.Printf("[!] Password Found!\n[!] SHA256: %s\n", password)
				}
			}
		}
	}
}
