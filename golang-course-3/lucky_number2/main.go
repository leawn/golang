package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
	usage    = `Welcome to the Lucky Number Game

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Wanna play?`
	randomMessage1 = `You were born as a winner`
	randomMessage2 = `Luck is clearly on your side. Well done!`
	randomMessage3 = `Sooo, you did that before, right? You win!`
)

func main() {
	messages := [3]string{randomMessage1, randomMessage2, randomMessage3}
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 && len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	for argsLoop := 0; argsLoop < len(args); argsLoop++ {
		guess, err := strconv.Atoi(args[argsLoop])
		if err != nil {
			fmt.Println("Not a number")
			return
		}

		if guess < 0 {
			fmt.Println("Please pick a positive number")
			return
		}

		for turn := 0; turn < maxTurns; turn++ {
			n := rand.Intn(guess + 1)

			if n == guess {
				m := rand.Intn(len(messages))
				fmt.Println(messages[m])
				if argsLoop == 0 || turn == 0 {
					fmt.Println("Wow, you are a lucker")
				}
				return
			}
		}
	}
	fmt.Println("You lost")
}
