package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Get the title and artist of the last played song on file
func loadLastPlayed() (string, string) {
	if _, err := os.Stat("lastPlayed.txt"); errors.Is(err, os.ErrNotExist) {
		return "", ""
	}

	file, err := os.Open("lastPlayed.txt")
	if err != nil {
		log.Println(err)
		return "", ""
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Println(err)
	}

	return text[0], text[1]
}

// Save the last played song to file
func saveLastPlayed(title string, artist string) error {
	text := []byte(fmt.Sprintf("%s\n%s", title, artist))
	err := os.WriteFile("lastPlayed.txt", text, 0644)
	if err != nil {
		return err
	}
	return nil
}
