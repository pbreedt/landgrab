package landgame

import (
	"log"
	"math/rand"
)

func (g *Gameboard) Initialize(players ...Player) {
	// areas := StakePlayerAreas(len(players))
	for i := 0; i < len(players); i++ {
		x, y := g.GetRandomPosition()
		g[x][y] = Block{Occupied_by: Entity{Etype: PlayerHome, Belongs_to: players[i]}}

		x, y = g.GetRandomPosition()
		g[x][y] = Block{Occupied_by: Entity{Etype: Rob}}

		x, y = g.GetRandomPosition()
		g[x][y] = Block{Occupied_by: Entity{Etype: Attack}}
	}
}

func (g Gameboard) GetRandomPosition() (int, int) {
	for {
		x := rand.Intn(boardsize)
		y := rand.Intn(boardsize)

		if g[x][y].Occupied_by.Etype == OpenSpace {
			return x, y
		} else {
			log.Default().Printf("%d,%d occupied by %s", x, y, g[x][y].Occupied_by)
		}
	}
}

type Area struct {
	startX, endX int
	startY, endY int
}

// boardsize = 1 bigger than index
func StakePlayerAreas(players int) []Area {
	areas := make([]Area, players-1)

	// assume boardsize of 12 x 12: (index 0,0 -> 11,11)
	// 2 players: 1) 0,0 -> 11,5  2) 0,6 -> 11,11
	if players == 2 {
		// top half
		areas[0].startX = 0
		areas[0].startY = 0
		areas[0].endX = boardsize - 1
		areas[0].endY = (boardsize / 2) - 1

		// bottom half
		areas[1].startX = 0
		areas[1].startY = (boardsize / 2)
		areas[1].endX = boardsize - 1
		areas[1].endY = boardsize - 1
	}

	// 3 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,3 -> 11,8
	if players == 3 {
		// top left quadrant
		areas[0].startX = 0
		areas[0].startY = 0
		areas[0].endX = (boardsize / 2) - 1
		areas[0].endY = (boardsize / 2) - 1

		// bottom left quadrant
		areas[1].startX = 0
		areas[1].startY = (boardsize / 2)
		areas[1].endX = (boardsize / 2) - 1
		areas[1].endY = boardsize - 1

		// top right quadrant, moved down by half a quadrant
		areas[2].startX = (boardsize / 2)
		areas[2].startY = (boardsize / 4)
		areas[2].endX = boardsize - 1
		areas[2].endY = boardsize - (boardsize / 4)
	}

	return areas
}

func (g *Gameboard) MarkArea(area Area) {
	for x := area.startX; x <= area.endX; x++ {
		for y := area.startY; x <= area.endY; x++ {
			g[x][y] = Block{}
		}
	}
}
