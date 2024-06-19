package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand generates a number between 0 and whatever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and no
	// greate than 10 (multiplying by 1 makes the game a bit silly)
	var firstNumber = rand.Intn(8) + 2
	var seconNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*seconNumber - subtraction
	playTheGame(firstNumber, seconNumber, subtraction, answer)

}
func playTheGame(firstNumber int, seconNumber int, subtraction int, answer int) {
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("________")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", seconNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the original thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	//gave them the answer

	fmt.Println("The answer is", answer)

}
