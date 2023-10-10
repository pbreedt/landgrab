package landgame

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/pbreedt/stdio/input"
)

func (g *Gameboard) Initialize(players ...Player) {
	g.players = append(g.players, players...)
	areas := StakePlayerAreas(len(players))

	for i := 0; i < len(players); i++ {
		// assign 1 Swap card to eack player
		g.players[i].SwapCards = append(g.players[i].SwapCards, SwapCard)

		// assign 1 Home block in player area
		x, y := g.GetRandomPosition(areas[i])
		g.board[y][x] = Block{Marker: HomeBlock.Marker, Belongs_to: &players[i]}
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

	RandomizeLandPieces(100)
	// visualize player areas
	// for i, a := range areas {
	// 	g.MarkArea(a, strconv.Itoa(i))
	// }
}

func (g *Gameboard) Play() {
	playerTurn := 0

	for {
		fmt.Println(g)

		fmt.Println(LandPieces.PrintN(5))

		move, _ := input.ReadString(fmt.Sprintf("Move for player %s? ", g.players[playerTurn].Name))
		// fmt.Printf("Move for player %s: %s", g.players[playerTurn].Name, move)

		//do move:
		switch strings.ToUpper(move) {
		case "Q":
			return
		default:
		}

		//next turn:
		playerTurn++
		if playerTurn >= len(g.players) {
			playerTurn = 0
		}
	}

}

func (g Gameboard) GetRandomPosition(area Area) (int, int) {
	for {
		x := rand.Intn(area.end.X-area.start.X) + area.start.X
		y := rand.Intn(area.end.Y-area.start.Y) + area.start.Y

		if g.board[y][x].Marker == OpenBlock.Marker {
			return x, y
		} else {
			log.Default().Printf("%d,%d occupied by %s", x, y, g.board[y][x].Belongs_to)
		}
	}
}

type Coordinate struct {
	X, Y int
}
type Area struct {
	start Coordinate
	end   Coordinate
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
		areas[0].start.X = 0
		areas[0].start.Y = 0
		areas[0].end.X = boardsize - 1
		areas[0].end.Y = (boardsize / 2) - 1

		// bottom half
		areas[1].start.X = 0
		areas[1].start.Y = (boardsize / 2)
		areas[1].end.X = boardsize - 1
		areas[1].end.Y = boardsize - 1
	}

	// 3 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,3 -> 11,8
	if players == 3 {
		// top left quadrant
		areas[0].start.X = 0
		areas[0].start.Y = 0
		areas[0].end.X = (boardsize / 2) - 1
		areas[0].end.Y = (boardsize / 2) - 1

		// bottom left quadrant
		areas[1].start.X = 0
		areas[1].start.Y = (boardsize / 2)
		areas[1].end.X = (boardsize / 2) - 1
		areas[1].end.Y = boardsize - 1

		// top right quadrant, moved down by half a quadrant
		areas[2].start.X = (boardsize / 2)
		areas[2].start.Y = (boardsize / 4)
		areas[2].end.X = boardsize - 1
		areas[2].end.Y = boardsize - (boardsize / 4) - 1
	}

	// 4 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,0 -> 11,5  4) 6,6 -> 11,11
	if players == 4 {
		// top left quadrant
		areas[0].start.X = 0
		areas[0].start.Y = 0
		areas[0].end.X = (boardsize / 2) - 1
		areas[0].end.Y = (boardsize / 2) - 1

		// bottom left quadrant
		areas[1].start.X = 0
		areas[1].start.Y = (boardsize / 2)
		areas[1].end.X = (boardsize / 2) - 1
		areas[1].end.Y = boardsize - 1

		// top right quadrant
		areas[2].start.X = (boardsize / 2)
		areas[2].start.Y = 0
		areas[2].end.X = boardsize - 1
		areas[2].end.Y = (boardsize / 2) - 1

		// bottom right quadrant
		areas[3].start.X = (boardsize / 2)
		areas[3].start.Y = (boardsize / 2)
		areas[3].end.X = boardsize - 1
		areas[3].end.Y = boardsize - 1
	}

	return areas
}

func (g *Gameboard) MarkArea(area Area, marker string) {
	b := Block{Marker: marker}

	for x := area.start.X; x <= area.end.X; x++ {
		for y := area.start.Y; y <= area.end.Y; y++ {
			g.board[y][x] = b
		}
	}
}
