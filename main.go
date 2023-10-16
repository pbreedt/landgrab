package main

import (
	"fmt"

	"github.com/pbreedt/landgrab/landgrab"
)

var (
	// Red     = landgrab.Player{Name: "Red", Color: landgrab.Red}
	Green   = landgrab.Player{Name: "Green", Color: landgrab.Green}
	Yellow  = landgrab.Player{Name: "Yellow", Color: landgrab.Yellow}
	Blue    = landgrab.Player{Name: "Blue", Color: landgrab.Blue}
	Magenta = landgrab.Player{Name: "Magenta", Color: landgrab.Magenta}
	Cyan    = landgrab.Player{Name: "Cyan", Color: landgrab.Cyan}
	White   = landgrab.Player{Name: "White", Color: landgrab.White}

	PlayerColors [7]landgrab.Player = [7]landgrab.Player{Green, Yellow, Blue, Magenta, Cyan, White}
)

func main() {
	// num_players, _ := input.ReadInt("How many players (2 / 3 / 4) ? ")
	num_players := 3
	if num_players < 2 || num_players > 4 {
		fmt.Println("Incorrect number of players selected.")
		return
	}

	var players []landgrab.Player
	gameboard := landgrab.NewGameboard()

	for i := 0; i < num_players; i++ {
		// name, _ := input.ReadString(fmt.Sprintf("Name for player %d? ", i))
		player := PlayerColors[i]
		player.Name = []string{"Jan", "Piet", "Koos"}[i]
		players = append(players, player)
	}

	gameboard.Initialize(players...)

	gameboard.Play()

	// num1 := uint16(rand.Intn(math.MaxUint16))
	// num2 := uint16(rand.Intn(math.MaxUint16))
	// num3 := uint16(rand.Intn(math.MaxUint16))
	// num4 := uint16(rand.Intn(math.MaxUint16))
	// num5 := uint16(rand.Intn(math.MaxUint16))

	// result := landgrab.ToBinaryGrid(num1, num2, num3, num4, num5)
	// fmt.Println(result)

}
