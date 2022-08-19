package main

import (
	"fmt"
	"math/rand"
	"time"
)

// computer programming is all about communication 
// your code communicates your intents to the compiler/computer 
// and when written well communicates your intentions to other people

func greet()  {
	fmt.Println("Hello world")
}

func printing() {
	fmt.Println("My weight on the surface of Mars is", 149*0.3783, "lbs, and i would be", 41*365.2425/687, "years old.")
	// the line of code above can be rewritten with Printf as
	fmt.Printf("My weight on the surface of Mars is %v lbs, and i would be %v years old", 149*0.3783, 41*365.2425/687)

	fmt.Printf("1*2 = %2v\n", 2)
	fmt.Printf("1*3 = %2v\n", 2)
	fmt.Printf("1*4 = %2v\n", 2)
}

func constantsAndVariables() {
	// constants and variables
	const lightspeed = 299792 // km/s
	var distance = 56000000 // km

	fmt.Printf("%v seconds\n", distance/lightspeed)

	distance = 401000000
	fmt.Printf("%v seconds\n", distance/lightspeed)

	// assignment operators shortcut

	const testConstant = 2

	var age = 40
	age += 3
	age += 5
	fmt.Println("My new age is", age)
}

func numbers() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
}

func conditionals() {
	var age = 12
	const VOTE_AGE = 18

	if age >= VOTE_AGE {
		fmt.Println("Congrats you are eligible to vote")
	} else {
		fmt.Println("You are not yet eligible to vote, please chill for the next", VOTE_AGE - age, "years")
	}
}

func repetition() {
	var count = 1
	for count <= 100 {
		var divisibleBy3 = count % 3 == 0
		var divisibleBy5 = count % 5 == 0
		switch {
		case divisibleBy3 && divisibleBy5:
			fmt.Println(count,"is divisible by 3 and 5")
		
		case divisibleBy3:
			fmt.Println(count,"is divisible by 3")
		
		case divisibleBy5:
			fmt.Println(count, "is divisible by 5")
		}
		count ++
	}
}

func infiniteRepetition() {
	var count = 1
	for {
		fmt.Printf("iterating at %v\n", count)
		if count == 100 {
			break;
		}
		count ++
	}
}

func guessGame() {
	var game = func () {
		rand.Seed(time.Now().UnixNano())
		var luckyNumber = rand.Intn(100)
		fmt.Println("\n\nWelcome to guess a number\n\n")
		for {
			var guess int
			fmt.Scanln(&guess)
			fmt.Println("your guess is", guess)
			if guess > luckyNumber {
				fmt.Println("your guess is too high")
			} else if guess < luckyNumber {
				fmt.Println("your guess is too low")
			} else if  guess == luckyNumber {
				fmt.Println("Congrats you guessed the number")
				break
			}
		}
		
		fmt.Println("Thats the end of the game, thank you for playing!")
	}

	for {
		game()

		fmt.Println("Press y/n to continue to stop playing")
		var decision string
		fmt.Scanln(&decision)

		if decision == "n" || decision == "no" {
			break
		} else if decision == "y" || decision == "yes" {
			continue;
		} else {
			fmt.Println("Invalid choice. existing game")
			break
		}
	}
}

// func scope() {
// 	var scopeVar = 2000
// 	func innerScope() {
// 		fmt.Print("scope variable", scopeVar)
// 	}
// }

func main() {
	// greet()
	// printing()
	// constantsAndVariables()
	// numbers()
	// conditionals()
	// repetition()
	// scope()
	// infiniteRepetition()
	guessGame()
}





