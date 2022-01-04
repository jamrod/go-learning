package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	playing := true
	for playing {
		target := rand.Intn(1000) + 1
		fmt.Println("I've chosen a number between 1 and 1000")
		fmt.Println("Can you guess it?")
		guessLeft := 10
		success := false
		reader := bufio.NewReader(os.Stdin)
		for guessLeft > 0 {
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)
			guess, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}
			if guess == target {
				fmt.Println("You guessed it!")
				success = true
				break
			} else if guess > target {
				fmt.Println("Lower...")
			} else {
				fmt.Println("Higher...")
			}
			guessLeft--
			fmt.Println(guessLeft, "Guesses remain")
		}
		if success {
			fmt.Println("You still had", guessLeft, "guesses remaining.")
		} else {
			fmt.Println("Sorry, you lost! The number was", target)
		}
		fmt.Println("Play again(y/n)?")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(input)
		input = strings.TrimSpace(input)
		fmt.Println(input == "y")

		if strings.ToLower(input) != "y" {
			fmt.Println("Bye!")
			playing = false
		}
	}
}
