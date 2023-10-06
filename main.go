package main

import (
	"fmt"
	"strconv"

	"github.com/pbreedt/landgame/landgame"
)

var (
	Red     = landgame.Player{Name: "Red", Color: landgame.Red}
	Green   = landgame.Player{Name: "Green", Color: landgame.Green}
	Yellow  = landgame.Player{Name: "Yellow", Color: landgame.Yellow}
	Blue    = landgame.Player{Name: "Blue", Color: landgame.Blue}
	Magenta = landgame.Player{Name: "Magenta", Color: landgame.Magenta}
	Cyan    = landgame.Player{Name: "Cyan", Color: landgame.Cyan}
	White   = landgame.Player{Name: "White", Color: landgame.White}

	PlayerColors [7]landgame.Player = [7]landgame.Player{Red, Green, Yellow, Blue, Magenta, Cyan, White}
)

func main() {
	// num_players, _ := input.ReadInt("How many players (2 / 3 / 4) ? ")
	num_players := 3
	if num_players < 2 || num_players > 4 {
		fmt.Println("Incorrect number of players selected.")
		return
	}

	var players []landgame.Player
	gameboard := landgame.NewGameboard()

	for i := 0; i < num_players; i++ {
		// name, _ := input.ReadString(fmt.Sprintf("Name for player %d? ", i))
		name := strconv.Itoa(i)
		player := PlayerColors[i]
		player.Name = name
		players = append(players, player)
	}

	gameboard.Initialize(players...)

	fmt.Println(gameboard)
}
