package main

import (
	"fmt"
)

//Player :Datos del usuario
type Player struct {
	Name    string
	Age     int
	Balance int64
	casino  string
}

func main() {

	fmt.Println("lecture 2")

	var players [2]Player

	players[0] = Player{"Pepito", 25, 30000, "zito"}
	fmt.Println(players[0])

	var onlyPlayer Player
	onlyPlayer.Name = "Maria"
	onlyPlayer.Age = 21
	onlyPlayer.Balance = 1000
	onlyPlayer.casino = "jester"
	players[1] = onlyPlayer
	fmt.Println(players[1])

}
