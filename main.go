package main

import (
	"fmt"

	"github.com/pbreedt/landgame/landgame"
	"github.com/pbreedt/stdio/input"
)

var (
	Red     = landgame.Player{Name: "Red", Color: landgame.Red}
	Green   = landgame.Player{Name: "Green", Color: landgame.Green}
	Yellow  = landgame.Player{Name: "Yellow", Color: landgame.Yellow}
	Blue    = landgame.Player{Name: "Blue", Color: landgame.Blue}
	Magenta = landgame.Player{Name: "Magenta", Color: landgame.Magenta}
	Cyan    = landgame.Player{Name: "Cyan", Color: landgame.Cyan}
	White   = landgame.Player{Name: "White", Color: landgame.White}

	Players [7]landgame.Player = [7]landgame.Player{Red, Green, Yellow, Blue, Magenta, Cyan, White}
)

func main() {
	num_players, _ := input.ReadInt("How many players (2 / 3 / 4) ? ")
	if num_players < 2 || num_players > 4 {
		fmt.Println("Incorrect number of players selected.")
		return
	}

	var players []landgame.Player
	gameboard := landgame.NewGameboard()

	for i := 0; i < num_players; i++ {
		player, _ := input.ReadString(fmt.Sprintf("Name for player %d? ", i))
		Players[i].Name = player
		players = append(players, Players[i])
	}

	// x, y := gameboard.GetRandomPosition()
	// gameboard[x][y] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.PlayerLand, Belongs_to: Red}}
	// // gameboard[2][5] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.Grass, Belongs_to: Green}}
	// x, y = gameboard.GetRandomPosition()
	// gameboard[x][y] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.PlayerLand, Belongs_to: Yellow}}
	// // gameboard[4][5] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.Grass, Belongs_to: Blue}}
	// x, y = gameboard.GetRandomPosition()
	// gameboard[x][y] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.PlayerLand, Belongs_to: Magenta}}
	// // gameboard[6][5] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.Grass, Belongs_to: Cyan}}
	// // gameboard[7][5] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.Grass, Belongs_to: White}}

	// for i := 0; i < 20; i++ {
	// 	x, y = gameboard.GetRandomPosition()
	// 	gameboard[x][y] = landgame.Block{Occupied_by: landgame.Entity{Etype: landgame.PlayerLand, Belongs_to: Magenta}}
	// }
	gameboard.Initialize(players...)

	fmt.Println(gameboard)
}
