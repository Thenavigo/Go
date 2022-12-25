package main

import (
	f "fmt"
	  "bufio"
	  "os"
	  "time"
	  "math/rand"
	  "strings"
	  "log"
	  "strconv"
)

var pl = f.Println

func main() {
	// Guess a number between 0 and 50 :
	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)	
	randNum := rand.Intn(50) + 1
	for true{
		f.Print(" Guess a number between 0 and 50 : ")
		pl("Random Number is : ", randNum)
		guess, err := reader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil{
			log.Fatal(err)
		}
		if iGuess > randNum{
			pl("Guess Something Lower")
		} else if iGuess < randNum {
			pl("Guess Something Higher")
		} else {
			pl("You Guessed it")
			break
		}
	}

}