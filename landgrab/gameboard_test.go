package landgrab

import (
	"testing"
)

var (
	PYellow  = Player{Name: "Yellow", Color: Yellow}
	PBlue    = Player{Name: "Blue", Color: Blue}
	PMagenta = Player{Name: "Magenta", Color: Magenta}
	PCyan    = Player{Name: "Cyan", Color: Cyan}

	PlayerColors [7]Player = [7]Player{PYellow, PMagenta, PCyan, PBlue}

	players []Player
)

func init() {
	for i := 0; i < 3; i++ {
		player := PlayerColors[i]
		player.Name = []string{"Jan", "Piet", "Koos"}[i]
		players = append(players, player)
	}
}

func getGameboard() Gameboard {
	gb := NewGameboard()

	gb.Initialize(players...)

	return gb
}

func TestMarkGameboard(t *testing.T) {
	gb := getGameboard()

	area := Area{
		Start: Coordinate{2, 4},
		End:   Coordinate{3, 7},
	}

	gb.MarkArea(area, "X")
	gb.Display()

	pass, got, pos := gb.checkAreaMarker(area, "X")
	if !pass {
		t.Fatalf("MarkArea failed. Expected marker 'X' at position %s, got '%s'", pos, got)
	}

	// test the test function :)
	// gb.Board[5][3].Marker = "o"
	// pass, got, pos = gb.checkAreaMarker(area, "X")
	// if !pass {
	// 	t.Fatalf("MarkArea failed. Expected marker 'X' at position %s, got '%s'", pos, got)
	// }
}

func (gb Gameboard) checkAreaMarker(area Area, marker string) (pass bool, found string, position Coordinate) {
	pass = true
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			if gb.Board[y][x].Marker != marker {
				pass = false
				found = gb.Board[y][x].Marker
				position = Coordinate{x, y}
				return
			}
		}
	}

	return
}

func TestGetMaxSqrArea(t *testing.T) {
	gb := getGameboard()

	area := Area{
		Start: Coordinate{0, 0},
		End:   Coordinate{5, 5},
	}

	gb.assignAreaToPlayer(area, &PYellow, LandPieceBlock.Marker)
	// gb.Display()

	// Remove LandPiece blocks
	gb.Board[2][0].Belongs_to = &Player{}
	gb.Board[2][0].Marker = ""
	gb.Board[0][1].Belongs_to = &Player{}
	gb.Board[0][1].Marker = ""

	gb.Display()

	maxArea, maxVal := gb.GetMaxSquareArea(area)
	expectMaxArea := Area{
		Start: Coordinate{1, 1},
		End:   Coordinate{5, 5},
	}

	if maxVal != 25 {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area value %d, got %d", 25, maxVal)
	}
	if maxArea.String() != expectMaxArea.String() {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area %s, got area %s", expectMaxArea, maxArea)
	}
}

func TestGetMaxSqrArea2(t *testing.T) {
	gb := getGameboard()

	area := Area{
		Start: Coordinate{0, 0},
		End:   Coordinate{5, 4},
	}

	gb.assignAreaToPlayer(area, &PYellow, LandPieceBlock.Marker)

	// Remove LandPiece blocks
	gb.Board[0][2].Belongs_to = &Player{}
	gb.Board[0][2].Marker = ""
	gb.Board[0][5].Belongs_to = &Player{}
	gb.Board[0][5].Marker = ""
	gb.Board[1][5].Belongs_to = &Player{}
	gb.Board[1][5].Marker = ""
	gb.Board[2][5].Belongs_to = &Player{}
	gb.Board[2][5].Marker = ""
	gb.Board[4][5].Belongs_to = &Player{}
	gb.Board[4][5].Marker = ""

	gb.Display()

	maxArea, maxVal := gb.GetMaxSquareArea(area)
	expectMaxArea := Area{
		Start: Coordinate{0, 1},
		End:   Coordinate{3, 4},
	}
	expectMaxVal := 16

	if maxVal != expectMaxVal {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area value %d, got %d", expectMaxVal, maxVal)
	}
	if maxArea.String() != expectMaxArea.String() {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area %s, got area %s", expectMaxArea, maxArea)
	}
}

func TestGetMaxSqrArea3(t *testing.T) {
	gb := getGameboard()

	area := Area{
		Start: Coordinate{0, 0},
		End:   Coordinate{3, 3},
	}

	gb.assignAreaToPlayer(area, &PYellow, LandPieceBlock.Marker)

	// Remove LandPiece blocks
	delarea := Area{
		Start: Coordinate{1, 1},
		End:   Coordinate{3, 3},
	}
	gb.assignAreaToPlayer(delarea, nil, "")

	gb.Display()

	maxArea, maxVal := gb.GetMaxSquareArea(area)
	expectMaxArea := Area{
		Start: Coordinate{0, 0},
		End:   Coordinate{0, 0},
	}
	expectMaxVal := 1

	if maxVal != expectMaxVal {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area value %d, got %d", expectMaxVal, maxVal)
	}
	if maxArea.String() != expectMaxArea.String() {
		t.Fatalf("GetMaxSolidArea failed. Expected max solid area %s, got area %s", expectMaxArea, maxArea)
	}
}

func (gb *Gameboard) assignAreaToPlayer(area Area, player *Player, marker string) {
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			gb.Board[y][x].Belongs_to = player
			gb.Board[y][x].Marker = marker
		}
	}
}
