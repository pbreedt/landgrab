package main

import (
	"fmt"

	"github.com/pbreedt/landgame/landgame"
)

func main() {
	gameboard := landgame.NewGameboard()

	Red := landgame.Player{Name: "Red", Color: landgame.Red}
	// Green := landgame.Player{Name: "Green", Color: landgame.Green}
	Yellow := landgame.Player{Name: "Yellow", Color: landgame.Yellow}
	// Blue := landgame.Player{Name: "Blue", Color: landgame.Blue}
	Magenta := landgame.Player{Name: "Magenta", Color: landgame.Magenta}
	// Cyan := landgame.Player{Name: "Cyan", Color: landgame.Cyan}
	// White := landgame.Player{Name: "White", Color: landgame.White}
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
	gameboard.Initialize(Red, Yellow, Magenta)

	fmt.Println(gameboard)
}
