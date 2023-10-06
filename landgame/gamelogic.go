package landgame

import (
	"log"
	"math/rand"
)

func (g *Gameboard) Initialize(players ...Player) {
	g.players = append(g.players, players...)
	areas := StakePlayerAreas(len(players))

	for i := 0; i < len(players); i++ {
		// assign 1 Home block in player area
		x, y := g.GetRandomPosition(areas[i])
		g.board[y][x] = Block{Marker: " H ", Belongs_to: &players[i]}
		// log.Default().Printf("done with home for %s", players[i].Name)

		// assign 1 Rob block in player area
		x, y = g.GetRandomPosition(areas[i])
		g.board[y][x] = RobBlock
		// log.Default().Printf("done with rob for %s", players[i].Name)

		// assign 1 Attack block in player area
		x, y = g.GetRandomPosition(areas[i])
		g.board[y][x] = AttackBlock
		// log.Default().Printf("done with attack for %s", players[i].Name)
	}

	// visualize player areas
	// for i, a := range areas {
	// 	g.MarkArea(a, strconv.Itoa(i))
	// }
}

func (g Gameboard) GetRandomPosition(area Area) (int, int) {
	for {
		x := rand.Intn(area.endX-area.startX) + area.startX
		y := rand.Intn(area.endY-area.startY) + area.startY

		if g.board[y][x].Marker == OpenBlock.Marker {
			return x, y
		} else {
			log.Default().Printf("%d,%d occupied by %s", x, y, g.board[y][x].Belongs_to)
		}
	}
}

type Area struct {
	startX, endX int
	startY, endY int
}

// StakePlayerAreas ensures each player has equal area
// home, rob & attack blocks are randomly placed inside each player's area
// (for simplicity's sake, only catering for 2 / 3 / 4 players for now)
// max index is 1 less than boardsize
func StakePlayerAreas(players int) []Area {
	areas := make([]Area, players)

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
		areas[2].endY = boardsize - (boardsize / 4) - 1
	}

	// 4 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,0 -> 11,5  4) 6,6 -> 11,11
	if players == 4 {
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

		// top right quadrant
		areas[2].startX = (boardsize / 2)
		areas[2].startY = 0
		areas[2].endX = boardsize - 1
		areas[2].endY = (boardsize / 2) - 1

		// bottom right quadrant
		areas[3].startX = (boardsize / 2)
		areas[3].startY = (boardsize / 2)
		areas[3].endX = boardsize - 1
		areas[3].endY = boardsize - 1
	}

	return areas
}

func (g *Gameboard) MarkArea(area Area, marker string) {
	b := Block{Marker: marker}

	for x := area.startX; x <= area.endX; x++ {
		for y := area.startY; y <= area.endY; y++ {
			g.board[y][x] = b
		}
	}
}
