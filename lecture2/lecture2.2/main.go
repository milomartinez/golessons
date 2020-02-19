package main

import (
	"fmt"

	"./players"
)

func main() {

	fmt.Println("lecture 2")

	player := players.New("Luchito", 12, 111, "zito")
	player.Print()

	fmt.Println("name:", player.Name)
	//fmt.Println("casino:", player.casino)
	fmt.Println("casino:", player.GetCasino())
}
