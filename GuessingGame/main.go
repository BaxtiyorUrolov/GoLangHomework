package main

import (
	"fmt"
	"math/rand"
)

func main() {
	name := getPlayerName()
	old := getPlayerAge()
	favouriteNumber := getPlayerFavouriteNumber()
	chances := getPlayerChances()

	player := NewPlayer(name, old, favouriteNumber, chances)
	game := Game{}

	for {
		game.NewGame(player)
		game.StartGame()
		if !playAgain() {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}

type Game struct {
	Player
	RandomNumber int
}

// Random number function

func (g *Game) NewGame(player Player) {
	if player.Chances > 0 && player.Chances < 5 {
		g.RandomNumber = rand.Intn(10)
	} else if player.Chances >= 5 && player.Chances < 10 {
		g.RandomNumber = rand.Intn(20)
	} else if player.Chances >= 10 && player.Chances < 15 {
		g.RandomNumber = rand.Intn(30)
	}
	g.Player = player
}

func (g *Game) StartGame() {
	fmt.Printf("Welcome %s\nYour age: %d\nYour favourite number: %d\n", g.Name, g.Old, g.FavouriteNumber)
	fmt.Println("This is a guessing game")
	for i := 0; i < g.Chances; i++ {
		var n int
		fmt.Printf("%d chance, enter number: ", i+1)
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if n == g.RandomNumber {
			if g.FavouriteNumber == g.RandomNumber {
				fmt.Println("Congratulations! You won by guessing your favorite number.")
			}

			fmt.Printf("You won! You found the random number in %d tries\n", i+1)
			return
		} else {
			fmt.Println("Incorrect")
		}
	}
	fmt.Println("You lost! The random number was", g.RandomNumber)
}

type Player struct {
	Name            string
	Old             int
	FavouriteNumber int
	Chances         int
}

func NewPlayer(name string, old, favouriteNumber, chances int) Player {
	return Player{
		Name:            name,
		Old:             old,
		FavouriteNumber: favouriteNumber,
		Chances:         chances,
	}
}

func getPlayerName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	return name
}

func getPlayerAge() int {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	return age
}

func getPlayerFavouriteNumber() int {
	var favouriteNumber int
	fmt.Print("Enter your favourite number: ")
	fmt.Scan(&favouriteNumber)
	return favouriteNumber
}

func getPlayerChances() int {
	var chances int
	fmt.Print("Enter your chances: ")
	fmt.Scan(&chances)
	return chances
}

func playAgain() bool {
	var response string
	fmt.Print("Do you want to play again? (yes/no): ")
	fmt.Scan(&response)
	return response == "yes"
}
