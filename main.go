package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var health1 int //player
var health2 int //cpu
var atk1 int
var atk2 int

func decreaseHP(res int) {
	if res == 0 {
		health1 = health1 - atk2
	}
	if res == 1 {
		health2 = health2 - atk1
	}
}

func battle(cpu string, player string) {
	if cpu == player {
		fmt.Print("Nothing happened")
		fmt.Println()
	}
	if cpu != player {
		//region cpu losecpu

		if cpu == "Rock" && (player == "p" || player == "paper") {
			health2 = health2 - atk1
			if health2 <= 0 {
				fmt.Print("You win")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			}
		} else if cpu == "Paper" && (player == "s" || player == "scissor") {
			health2 = health2 - atk1
			if health2 <= 0 {
				fmt.Print("You win")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			}
		} else if cpu == "Scissor" && (player == "r" || player == "rock") {
			health2 = health2 - atk1
			if health2 <= 0 {
				fmt.Print("You win")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			} //endregion
		} else if cpu == "Scissor" && (player == "p" || player == "paper") {
			health1 = health1 - atk2
			if health1 <= 0 {
				fmt.Print("You lose")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			}
		} else if cpu == "Paper" && (player == "r" || player == "rock") {
			health1 = health1 - atk2
			if health1 <= 0 {
				fmt.Print("You lose")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			}
		} else if cpu == "Rock" && (player == "s" || player == "scissor") {
			health1 = health1 - atk2
			if health1 <= 0 {
				fmt.Print("You lose")
				fmt.Scan()
				fmt.Println()
				os.Exit(1)
			}
		}

	}
}
func main() {
	var choose string
	choose2 := []string{"Rock", "Paper", "Scissor"}

	fmt.Println("WELCOME TO RPS FIGHT")
	fmt.Println("The Rules is You HAVE TO DEFEAT YOUR OPPONENT using ROCK PAPER SCISSOR")
	fmt.Println("Rocks win to scissor,Scissor win to paper, Paper win to Rock")
	fmt.Println("You and your opponent will get random health point and random attack point")
	fmt.Println("Good Luck")
	rand.Seed(time.Now().UTC().UnixNano())
	health1 = rand.Intn(100) + 1
	atk1 = rand.Intn(30) + 1
	health2 = rand.Intn(100) + 1
	atk2 = rand.Intn(30) + 1

	fmt.Println("You have " + strconv.Itoa(health1) + "HP")
	fmt.Println("You have " + strconv.Itoa(atk1) + "ATK")

	fmt.Println("CPU have " + strconv.Itoa(health2) + "HP")
	fmt.Println("CPU have " + strconv.Itoa(atk2) + "ATK")
	turn := 1
	choosen := choose2[rand.Intn(len(choose2))]

	fmt.Println(choosen)
	for {
		if health1 != 0 || health2 != 0 {
			fmt.Println("Battle Turn :" + strconv.Itoa(turn))
			fmt.Println("CPU Health  " + strconv.Itoa(health2))
			fmt.Println("Player Health  " + strconv.Itoa(health1))
			fmt.Print("Choose Rock(r)|Paper(p)|Scissor(s) :")
			fmt.Scan(&choose)
			//translate Player choose
			choices := strings.ToLower(choose)
			if choices == "rock" || choices == "r" || choices == "paper" || choices == "p" || choices == "scissor" || choices == "s" {
				fmt.Println("Choices has been locked")
			} else {
				fmt.Println("Please input between these(Rock|r|Paper|p|Scissor|s")
				continue
			}
			// randomIndex = rand.Intn(len(choose2))
			// fmt.Print(randomIndex)
			// pick := choose2[randomIndex]
			choosen = choose2[rand.Intn(len(choose2))]

			fmt.Println("You choose " + choose)
			fmt.Println("CPU choose " + choosen)
			battle(choosen, choose)

		}
	}
}
