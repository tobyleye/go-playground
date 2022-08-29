package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
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
	const maxGuess = 5

	var game = func () {
		rand.Seed(time.Now().UnixNano())
		var luckyNumber = rand.Intn(100)
		var nguesses = 0


		for {
			var guess int
			fmt.Scanln(&guess)
			nguesses++
			if guess > luckyNumber {
				fmt.Printf("Trial %v: your guess is too high", nguesses)
			} else if guess < luckyNumber {
				fmt.Printf("Trial %v: your guess is too low", nguesses)
			} else if  guess == luckyNumber {
				fmt.Printf("Trial %v: Congrats you guessed the number", nguesses)
				break
			}	

			fmt.Println()

			if nguesses == maxGuess {
				fmt.Println("You have exeeded your max no. of guesses. The answer is", luckyNumber)
				break;
			}

		}
		
		fmt.Println("Thats the end of the game, thank you for playing!")
	}	

	fmt.Println("\n\nWelcome to guess a number, the game!\n\n")

	for {
		// start game
		game()
		fmt.Println("Press y/n to continue or stop playing")
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

func scope() {
	// var scopeVar = 2000

	// short declaration 
	// scopeVar := 2000
	// the short declaration style is more concise and save some typing
	// it can also be used where the var declaration can't be used.

	for count := 100; count > 0; count-- {
		fmt.Println("count = ", count)
	}

	if condition := true; condition == true {
		fmt.Println("condition is true")
	} else {
		fmt.Println("condition is false")
	}

	// fmt.Println("final count is ", count)
	
}

var era =  "AD"

func scopeExample() {
	year := rand.Intn(100) + 2000 // random year between 2000 & 2100
	month := rand.Intn(12)
	var daysInMonth int;
	isLeapYear := year % 4 == 0

	switch month  {
	case 2:
		if isLeapYear {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4,6,9,11:
		daysInMonth = 30
	default:
		daysInMonth = 31
	}
	day := rand.Intn(daysInMonth) +1
	fmt.Println(era, year, month, day)

	// generally any opening curly braces introduces a new scope that ends 
	// with a closing braces. Any variable declared directly in the scope 
	// will be accessible throughout the span of that scope

	// case & default also introduces scope although they do no not use braces
}


func ticketToMars() {
	// ticket to mars

	speedMin := 16
	speedMax := 30
	// distance in km
	distance := 62100000
	priceMin := 36
	priceMax := 50
	// departureDate := "October 13, 2020"

	fmt.Printf("%-20v %-12v %-16v %4v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println(strings.Repeat("=", 56))

	for i:=0; i<10; i++ {
		rand.Seed(time.Now().UnixNano())
		var spaceline string

		speed := speedMin + rand.Intn(speedMax-speedMin)
		duration := distance / speed

		durationInDays := duration / (60 * 60 * 24)	

		price := priceMin + rand.Intn(priceMax-priceMin)

		var tripType string
		
		if trip := rand.Float32(); trip <= 0.5 {
			tripType = "One-way"
		} else if trip > 0.5 {
			tripType = "Round-trip"
			price *= 2
		}

		switch randomSpaceline := rand.Intn(3); randomSpaceline {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		fmt.Printf("%-20v %-12v %-16v $%3v\n", spaceline, durationInDays, tripType, price)

	}
	
}

func formats() {
	fmt.Printf("%-20v $%6v\n", "Oyeleye Oluwatobi", 52500)
}


func loops() {
	i := 0
	for {
		fmt.Println("hello world", i)
		i++

		if i == 5 {
			break;
		}
	}
}

func main() {
	// greet()
	// printing()
	// constantsAndVariables()
	// numbers()
	// conditionals()
	// repetition()
	// infiniteRepetition()
	// loops()
	// guessGame()
	// scope()
	// scopeExample()
	// formats()
	ticketToMars()
}





