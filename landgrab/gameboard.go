package landgrab

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

const boardsize int = 20

type Gameboard struct {
	Board              [boardsize][boardsize]Block
	Players            []Player
	LandPieces         *LandPiecesSlice
	currentPieceIndex  int
	currentPlayerIndex int
	showNumPieces      int
}

// Intended for initial configuration of gameboard
func NewGameboard() Gameboard {
	var gb Gameboard

	return gb
}

func (gb *Gameboard) CurrentPlayer() *Player {
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb *Gameboard) NextPlayer() *Player {
	gb.currentPlayerIndex++
	if gb.currentPlayerIndex >= len(gb.Players) {
		gb.currentPlayerIndex = 0
	}
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb *Gameboard) Display() *LandPiece {
	fmt.Printf("cur.piece:%d, cur.player.idx:%d, cur.player:%s\n", gb.currentPieceIndex, gb.currentPlayerIndex, gb.CurrentPlayer().Name)
	fmt.Println(gb)
	curPcIdx := gb.LandPieces.PrintUnplacedN(gb.currentPieceIndex, gb.showNumPieces)
	return &(*gb.LandPieces)[curPcIdx]
}

func (gb *Gameboard) String() string {
	bsh := ""
	for i, p := range gb.Players {
		bsh += fmt.Sprintf("Player %d: %v\n", i+1, p)
	}
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			bs += fmt.Sprintf("[%+v]", gb.Board[y][x])
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Marker         string
	Belongs_to     *Player
	HighlightErr   bool
	HighlightOK    bool
	LandPieceValue uint16
}

func (b Block) String() string {
	if b.HighlightErr {
		return ColorString(fmt.Sprintf("%1s", b.Marker), Red)
	}
	if b.HighlightOK {
		return ColorString(fmt.Sprintf("%1s", b.Marker), Green)
	}
	if b.Belongs_to != nil {
		return ColorString(fmt.Sprintf("%1s", b.Marker), b.Belongs_to.Color)
	}
	return fmt.Sprintf("%1s", b.Marker)
}

var (
	OpenBlock      Block = Block{Marker: ""}
	SwapBlock      Block = Block{Marker: "S"}
	GrabBlock      Block = Block{Marker: "G"}
	RockBlock      Block = Block{Marker: "R"}
	HomeBlock      Block = Block{Marker: "H"}
	LandPieceBlock Block = Block{Marker: "#"}
)

type Color int

const (
	Black  Color = iota // unused
	White               // default color
	Red                 // highlight: error
	Green               // highlight: ok
	Yellow              // player colors
	Blue
	Magenta
	Cyan
)

func (c Color) String() string {
	return [...]string{"Black", "White", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan"}[c]
}

func ColorString(s string, c Color) string {
	// log.Default().Printf("print '%s' in %s", s, c)
	switch c {
	case Red:
		return color.RedString(s)
	case Green:
		return color.GreenString(s)
	case Yellow:
		return color.YellowString(s)
	case Blue:
		return color.BlueString(s)
	case Magenta:
		return color.MagentaString(s)
	case Cyan:
		return color.CyanString(s)
	case White:
		return color.WhiteString(s)
	}
	return s
}

type Card int

const (
	SwapCard Card = iota
	GrabCard
	RockCard
)

func (c Card) String() string {
	return [...]string{"Swap", "Grab", "Rock"}[c]
}

type Player struct {
	Name      string
	Color     Color
	SwapCards []Card
	GrabCards []Card
	RockCards []Card
}

func (p Player) String() string {
	return ColorString(fmt.Sprintf("%10s | Swap Cards:%d, Grab Cards:%d, Rock Cards:%d", p.Name, len(p.SwapCards), len(p.GrabCards), len(p.RockCards)), p.Color)
}

func (gb *Gameboard) RemoveLandPieceFromGameboard(landPieceIdx int) {
	if (*gb.LandPieces)[landPieceIdx].PlacedAt == nil {
		return
	}

	lp := (*gb.LandPieces)[landPieceIdx]

	lpStr := lp.String()

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			sidx := x + (y * 4)
			if lpStr[sidx:sidx+1] == "#" && (lp.PlacedAt.Y+y >= 0) && (lp.PlacedAt.X+x >= 0) && (lp.PlacedAt.Y+y < (boardsize + 4)) && (lp.PlacedAt.X+x < (boardsize + 4)) {
				gb.Board[lp.PlacedAt.Y+y][lp.PlacedAt.X+x] = Block{}
			}
		}
	}

	(*gb.LandPieces)[landPieceIdx].PlacedAt = nil
	// log.Default().Printf("Land piece:%s mem:%p, placed:%v\n", (*gb.LandPieces)[landPieceIdx], &(*gb.LandPieces)[landPieceIdx], (*gb.LandPieces)[landPieceIdx].PlacedAt)
}

// Move LandPiece from used index to just before current piece index
// then set current piece index 1 back
func (gb *Gameboard) SetNextLandPieceTo(landPieceIdx int) {
	if landPieceIdx < 0 || landPieceIdx >= len(*gb.LandPieces) {
		return
	}

	removedLandPiece := (*gb.LandPieces)[landPieceIdx]
	// Remove LandPiece from slice
	(*gb.LandPieces) = append((*gb.LandPieces)[:landPieceIdx], (*gb.LandPieces)[landPieceIdx+1:]...)

	// Add LandPiece back at index before current piece index (since slice is shorter now)
	(*gb.LandPieces) = append((*gb.LandPieces)[:gb.currentPieceIndex-1], append(LandPiecesSlice{removedLandPiece}, (*gb.LandPieces)[gb.currentPieceIndex-1:]...)...)

	gb.currentPieceIndex--
}

// MarkArea places provided marker in all blocks of the provided area
// For testing only
func (gb *Gameboard) MarkArea(area Area, marker string) {
	b := Block{Marker: marker}

	log.Default().Printf("Marking area: %s with marker %s\n", area, marker)
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			gb.Board[y][x] = b
		}
	}
}

func (gb *Gameboard) ReplacePlayer(oldPlayer *Player, newPlayer *Player) {
	log.Default().Printf("Replacing player: %s with %s\n", oldPlayer.Name, newPlayer.Name)
	for x := 0; x < boardsize; x++ {
		for y := 0; y < boardsize; y++ {
			if gb.Board[y][x].Belongs_to != nil && gb.Board[y][x].Belongs_to.Name == oldPlayer.Name {
				gb.Board[y][x].Belongs_to = newPlayer
			}
		}
	}
}

func (gb *Gameboard) IsTouchingOwn(c Coordinate, p Player) bool {
	// This logic only allows 'direct' touching (immediately left/right or above/below)
	testCoords := []Coordinate{
		{X: c.X - 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y},
		{X: c.X, Y: c.Y - 1},
		{X: c.X, Y: c.Y + 1},
	}

	for _, tstCoord := range testCoords {
		if tstCoord.X >= 0 && tstCoord.X < boardsize && tstCoord.Y >= 0 && tstCoord.Y < boardsize {
			if gb.Board[tstCoord.Y][tstCoord.X].Belongs_to != nil && gb.Board[tstCoord.Y][tstCoord.X].Belongs_to.Name == p.Name {
				return true
			}
		}
	}

	// This logic includes 'diagonal' touching
	// area := Area{Start: Coordinate{X: c.X - 1, Y: c.Y - 1}, End: Coordinate{X: c.X + 1, Y: c.Y + 1}}

	// for x := area.Start.X; x <= area.End.X; x++ {
	// 	if x >= 0 && x < boardsize {
	// 		for y := area.Start.Y; y <= area.End.Y; y++ {
	// 			if y >= 0 && y < boardsize {
	// 				if gb.Board[y][x].Belongs_to != nil && gb.Board[y][x].Belongs_to.Name == p.Name {
	// 					return true
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return false
}
