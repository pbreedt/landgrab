package landgrab

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/pbreedt/stdio/input"
)

func (g *Gameboard) Initialize(players ...Player) {
	g.Players = append(g.Players, players...)
	areas := StakePlayerAreas(len(players))

	for i := 0; i < len(players); i++ {
		// assign 1 Swap card to eack player
		g.Players[i].SwapCards = append(g.Players[i].SwapCards, SwapCard)

		// assign 1 Home block in player area
		x, y := g.GetRandomPosition(areas[i])
		g.Board[y][x] = Block{Marker: HomeBlock.Marker, Belongs_to: &players[i]}
		// log.Default().Printf("done with home for %s", players[i].Name)

		// assign 1 Rob block in player area
		x, y = g.GetRandomPosition(areas[i])
		g.Board[y][x] = RobBlock
		// log.Default().Printf("done with rob for %s", players[i].Name)

		// assign 1 Attack block in player area
		x, y = g.GetRandomPosition(areas[i])
		g.Board[y][x] = AttackBlock
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
		keepTurn := false

		fmt.Println(g)

		fmt.Println(LandPieces.PrintN(5))

		move, _ := input.ReadString(fmt.Sprintf("Move for player %s? [Q]uit | [P]lace a piece ", g.Players[playerTurn].Name))
		// move := "P"
		// fmt.Printf("Move for player %s: %s", g.players[playerTurn].Name, move)

		//do move:
		switch strings.ToUpper(move) {
		case "Q": // Quit
			return
		case "P": // Place a piece
			placeXY := Coordinate{X: 0, Y: 0}
			keepPlacing := true
			itFits, gb := g.PlacePiece(&g.Players[playerTurn], LandPieces[0], placeXY)
			var moves []string
			for keepPlacing {
				if itFits {
					moves = []string{"[P]lace", "Move [R]ight", "Move [L]eft", "Move [U]p", "Move [D]own"}
				}
				placeMove, _ := input.ReadString(strings.Join(moves, " | ") + " ")

				switch strings.ToUpper(placeMove) {
				case "P":
					g = &gb
					keepPlacing = false
				case "R":
					placeXY.X++
					itFits, gb = g.PlacePiece(&g.Players[playerTurn], LandPieces[0], placeXY)
					if !itFits {
						placeXY.X--
						moves = []string{"[P]lace", "Move [L]eft", "Move [U]p", "Move [D]own"}
					}
				case "L":
					placeXY.X--
					itFits, gb = g.PlacePiece(&g.Players[playerTurn], LandPieces[0], placeXY)
					if !itFits {
						placeXY.X++
						moves = []string{"[P]lace", "Move [R]ight", "Move [U]p", "Move [D]own"}
					}
				case "U":
					placeXY.Y--
					itFits, gb = g.PlacePiece(&g.Players[playerTurn], LandPieces[0], placeXY)
					if !itFits {
						placeXY.Y++
						moves = []string{"[P]lace", "Move [R]ight", "Move [L]eft", "Move [D]own"}
					}
				case "D":
					placeXY.Y++
					itFits, gb = g.PlacePiece(&g.Players[playerTurn], LandPieces[0], placeXY)
					if !itFits {
						placeXY.Y--
						moves = []string{"[P]lace", "Move [R]ight", "Move [L]eft", "Move [U]p"}
					}
				}
			}
		default:
			keepTurn = true
		}

		if !keepTurn {
			//next turn:
			playerTurn++
			if playerTurn >= len(g.Players) {
				playerTurn = 0
			}
		}
	}

}

// PlacePiece checks if a LandPiece fits on Gameboard at provided Coordinate
// Then renders the Gameboard depicting the new LandPiece at the provided Coordinate
// Returns confirmation if the piece could fit & the new Gameboard
func (g Gameboard) PlacePiece(p *Player, lp LandPiece, c Coordinate) (bool, Gameboard) {
	binStr := lp.String()

	// TODO: replace below with g.LandPieceFits()
	if c.Y < 0 || c.Y > 8 || c.X < 0 || c.X > 8 {
		return false, g
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			sidx := x + (y * 4)
			if binStr[sidx:sidx+1] == "#" && (c.Y+y >= 0) && (c.X+x >= 0) && (c.Y+y < 16) && (c.X+x < 16) {
				g.Board[c.Y+y][c.X+x] = Block{Marker: "#", Belongs_to: p}
			}
		}
	}

	fmt.Println(g)

	return true, g // it fits
}

func (g Gameboard) GetRandomPosition(area Area) (int, int) {
	for {
		x := rand.Intn(area.End.X-area.Start.X) + area.Start.X
		y := rand.Intn(area.End.Y-area.Start.Y) + area.Start.Y

		if g.Board[y][x].Marker == OpenBlock.Marker {
			return x, y
		} else {
			log.Default().Printf("%d,%d occupied by %s", x, y, g.Board[y][x].Belongs_to)
		}
	}
}

type Coordinate struct {
	X, Y int
}
type Area struct {
	Start Coordinate
	End   Coordinate
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
		areas[0].Start.X = 0
		areas[0].Start.Y = 0
		areas[0].End.X = boardsize - 1
		areas[0].End.Y = (boardsize / 2) - 1

		// bottom half
		areas[1].Start.X = 0
		areas[1].Start.Y = (boardsize / 2)
		areas[1].End.X = boardsize - 1
		areas[1].End.Y = boardsize - 1
	}

	// 3 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,3 -> 11,8
	if players == 3 {
		// top left quadrant
		areas[0].Start.X = 0
		areas[0].Start.Y = 0
		areas[0].End.X = (boardsize / 2) - 1
		areas[0].End.Y = (boardsize / 2) - 1

		// bottom left quadrant
		areas[1].Start.X = 0
		areas[1].Start.Y = (boardsize / 2)
		areas[1].End.X = (boardsize / 2) - 1
		areas[1].End.Y = boardsize - 1

		// top right quadrant, moved down by half a quadrant
		areas[2].Start.X = (boardsize / 2)
		areas[2].Start.Y = (boardsize / 4)
		areas[2].End.X = boardsize - 1
		areas[2].End.Y = boardsize - (boardsize / 4) - 1
	}

	// 4 players: 1) 0,0 -> 5,5  2) 0,6 -> 5,11  3) 6,0 -> 11,5  4) 6,6 -> 11,11
	if players == 4 {
		// top left quadrant
		areas[0].Start.X = 0
		areas[0].Start.Y = 0
		areas[0].End.X = (boardsize / 2) - 1
		areas[0].End.Y = (boardsize / 2) - 1

		// bottom left quadrant
		areas[1].Start.X = 0
		areas[1].Start.Y = (boardsize / 2)
		areas[1].End.X = (boardsize / 2) - 1
		areas[1].End.Y = boardsize - 1

		// top right quadrant
		areas[2].Start.X = (boardsize / 2)
		areas[2].Start.Y = 0
		areas[2].End.X = boardsize - 1
		areas[2].End.Y = (boardsize / 2) - 1

		// bottom right quadrant
		areas[3].Start.X = (boardsize / 2)
		areas[3].Start.Y = (boardsize / 2)
		areas[3].End.X = boardsize - 1
		areas[3].End.Y = boardsize - 1
	}

	return areas
}

func (g *Gameboard) MarkArea(area Area, marker string) {
	b := Block{Marker: marker}

	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			g.Board[y][x] = b
		}
	}
}
