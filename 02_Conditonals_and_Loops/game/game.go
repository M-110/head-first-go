// this game challenges a player to guess a random number.
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
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1
	tries := 0
	for tries < 10 {
		word := "tries"
		if tries == 9 {
			word = "try"
		}
		fmt.Println(10-tries, word, "remaining")
		fmt.Print("Guess a number between 1 and 100: ")
		tries++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == answer {
			break
		} else if guess < answer {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}

	if tries == 10 {
		fmt.Println("You lose ...😥 The correcet answer was", answer)
	} else {
		fmt.Println("You win! 🤩")
	}
}
