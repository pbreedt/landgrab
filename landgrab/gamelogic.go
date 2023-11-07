package landgrab

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type GameSummary struct {
	FinalBoard    Gameboard
	PlayerSummary []PlayerSummary
}
type PlayerSummary struct {
	Player            Player
	BiggestArea       Area
	BiggestSquareArea Area
}

func (gb *Gameboard) Initialize(players ...Player) {
	gb.Players = append(gb.Players, players...)
	areas := StakePlayerAreas(len(players))
	gb.showNumPieces = 5

	for i := 0; i < len(players); i++ {
		// assign 1 Swap card to eack player
		gb.Players[i].SwapCards = append(gb.Players[i].SwapCards, SwapCard)

		// assign 1 Home block in player area
		x, y := gb.GetRandomPosition(areas[i])
		gb.Board[y][x] = Block{Marker: HomeBlock.Marker, Belongs_to: &players[i]}
		// log.Default().Printf("done with home for %s", players[i].Name)

		// assign 1 Rob block in player area
		x, y = gb.GetRandomPosition(areas[i])
		gb.Board[y][x] = GrabBlock
		// log.Default().Printf("done with rob for %s", players[i].Name)

		// assign 1 Attack block in player area
		x, y = gb.GetRandomPosition(areas[i])
		gb.Board[y][x] = RockBlock
		// log.Default().Printf("done with attack for %s", players[i].Name)
	}

	gb.LandPieces = RandomizeLandPieces(100)
	// visualize player areas
	// for i, a := range areas {
	// 	g.MarkArea(a, strconv.Itoa(i))
	// }
}

func (gb *Gameboard) Play() {
	done := false

	for !done {
		keepTurn := false

		curLandPiece := gb.Display()
		mainMenu := Menu{Category: "Action", Options: []Option{
			{Display: "[P]lace a piece", ActionKey: "P"},
			{Display: "[Q]uit", ActionKey: "Q"},
		}}
		cardMenu := GetCardMenu(*gb.CurrentPlayer())
		move := GetPlayerMove(*gb.CurrentPlayer(), mainMenu, cardMenu)

		// TESTING
		// move := "P"
		// fmt.Printf("Move for player %s: %s", g.players[playerTurn].Name, move)
		if !IsValidMove(move, mainMenu, cardMenu) {
			log.Default().Printf("Invalid option: %s\n", move)
			keepTurn = true
			continue
		}

		//do move:
		switch strings.ToUpper(move) {
		case "Q": // Quit
			done = true
		case "G": // Grab
			lp := gb.GrabPiece()
			if lp != nil {
				gb = gb.PlacePiece(lp)
			} else {
				keepTurn = true
			}
		case "S": // Swap
			gb.SwapPiece()
			keepTurn = true
		case "R": // Rock
			keepTurn = true
			gb.PlaceRock()
		case "P": // Place a piece
			gb = gb.PlacePiece(curLandPiece)
		default:
			keepTurn = true
		}

		if !keepTurn {
			//next turn:
			gb.NextPlayer()
		}

		if !gb.CheckOpenSpace() {
			done = true
		}
	}

	gs := gb.CheckOccupiedSpace()
	gs.FinalBoard.Display()

	best := 0
	var winner string

	for _, ps := range gs.PlayerSummary {
		if ps.BiggestSquareArea.Size() >= best {
			if best > 0 && ps.BiggestSquareArea.Size() == best {
				// draw
				winner = strings.ReplaceAll(winner, "THE WINNERS ARE ", "THE WINNERS ARE "+ColorString(strings.ToUpper(ps.Player.Name), ps.Player.Color)+" & ")
				winner = strings.ReplaceAll(winner, "THE WINNER IS ", "THE WINNERS ARE "+ColorString(strings.ToUpper(ps.Player.Name), ps.Player.Color)+" & ")
			} else {
				best = ps.BiggestSquareArea.Size()
				winner = "THE WINNER IS " + ColorString(strings.ToUpper(ps.Player.Name), ps.Player.Color) + "!!"
			}
		}
	}

	fmt.Println(winner)
}

// CheckOpenSpace returns true/false to indicate if enough open space remains to place more LandPieces
func (gb Gameboard) CheckOpenSpace() bool {
	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			block := gb.Board[y][x]
			if block.Belongs_to == nil || block.Marker == OpenBlock.Marker {
				return true
			}
		}
	}
	return false
}

func (gb Gameboard) CheckOccupiedSpace() GameSummary {
	gs := GameSummary{PlayerSummary: make([]PlayerSummary, len(gb.Players))}

	playerArea := make([]*Area, len(gb.Players))

	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			if gb.Board[y][x].Belongs_to != nil {
				for i, p := range gb.Players {
					if gb.Board[y][x].Belongs_to.Name == p.Name {
						if playerArea[i] == nil {
							playerArea[i] = &Area{Start: Coordinate{X: 11, Y: 11}}
						}
						if x < playerArea[i].Start.X {
							playerArea[i].Start.X = x
						}
						if x > playerArea[i].End.X {
							playerArea[i].End.X = x
						}
						if y < playerArea[i].Start.Y {
							playerArea[i].Start.Y = y
						}
						if y > playerArea[i].End.Y {
							playerArea[i].End.Y = y
						}
					}
				}
			}
		}
	}

	mostArea := 0
	for i, a := range playerArea {

		maxArea, maxVal := gb.GetMaxSquareArea(*a)
		if maxVal > mostArea {
			mostArea = maxVal
		}

		gs.PlayerSummary[i] = PlayerSummary{
			Player:            gb.Players[i],
			BiggestArea:       *a,
			BiggestSquareArea: maxArea,
		}

		log.Default().Printf("Player %s max solid area: %s %d", gb.Players[i].Name, maxArea, maxVal)
		gb.MarkArea(maxArea, ColorString(gb.Players[i].Name[0:1], gb.Players[i].Color))
	}

	gs.FinalBoard = gb

	return gs
}

func (gb Gameboard) GetMaxSquareArea(area Area) (maxArea Area, areaValue int) {
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			start := Coordinate{x, y}
			end := Coordinate{x, y}
			for end.X < 12 && end.X <= area.End.X && end.Y < 12 && end.Y <= area.End.Y {
				tstArea := Area{Start: start, End: end}
				solid, a := gb.IsSolidArea(tstArea)
				if solid && a > areaValue {
					areaValue = a
					maxArea = tstArea
					log.Default().Println("TEMP Max area:", areaValue, maxArea)
				}
				end.X++
				end.Y++
			}
		}
	}
	log.Default().Println("MAX Max area:", areaValue, maxArea)
	return
}

func (gb Gameboard) IsSolidArea(area Area) (bool, int) {
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			// log.Default().Printf("Solid test: x=%d, y=%d\n", x, y)
			if gb.Board[y][x].Marker != LandPieceBlock.Marker && gb.Board[y][x].Marker != HomeBlock.Marker {
				return false, 0
			}
		}
	}

	return true, area.Size() //((area.End.X - area.Start.X) + 1) * ((area.End.Y - area.Start.Y) + 1)
}

func (gb *Gameboard) PlaceRock() {
	placePos := Coordinate{0, 0}
	ok, ngb := gb.PreviewPlaceRock(gb.CurrentPlayer(), placePos)

	for {
		selectMenu := GetMovePieceMenu("Move rock")
		selectMenu.Options = append(selectMenu.Options, Option{Display: "[C]ancel place", ActionKey: "C"})
		if ok {
			selectMenu.Options = append(selectMenu.Options, Option{Display: "[P]lace rock", ActionKey: "P"})
		}

		ngb.Display()
		pieceMove := GetPlayerMove(*gb.CurrentPlayer(), selectMenu)
		if IsValidMove(pieceMove, selectMenu) {
			switch strings.ToUpper(pieceMove) {
			case "C": // Don't place rock
				log.Default().Printf("No rock placed")
				return
			case "P": // Place rock
				log.Default().Printf("Rock placed at %s\n", placePos)
				// Remove Player's Grab card
				gb.CurrentPlayer().RockCards = gb.CurrentPlayer().RockCards[1:]
				gb.UpdateCards()
				gb.Board[placePos.Y][placePos.X] = Block{Marker: RockBlock.Marker, Belongs_to: gb.CurrentPlayer()}
				return
			case "R": // Right
				if (placePos.X + 1) < 12 {
					placePos.X++
				} else {
					selectMenu.RemoveOption("R")
				}
				ok, ngb = gb.PreviewPlaceRock(gb.CurrentPlayer(), placePos)
			case "L": // Left
				if (placePos.X - 1) >= 0 {
					placePos.X--
				} else {
					selectMenu.RemoveOption("L")
				}
				ok, ngb = gb.PreviewPlaceRock(gb.CurrentPlayer(), placePos)
			case "U": // Up
				if (placePos.Y - 1) >= 0 {
					placePos.Y--
				} else {
					selectMenu.RemoveOption("D")
				}
				ok, ngb = gb.PreviewPlaceRock(gb.CurrentPlayer(), placePos)
			case "D": // Down
				if (placePos.Y + 1) < 12 {
					placePos.Y++
				} else {
					selectMenu.RemoveOption("D")
				}
				ok, ngb = gb.PreviewPlaceRock(gb.CurrentPlayer(), placePos)
			}
		}
	}
}

func (gb Gameboard) PreviewPlaceRock(p *Player, c Coordinate) (bool, Gameboard) {
	placeOK := false

	// cannot select at coordinate
	if c.Y < 0 || c.Y >= 12 || c.X < 0 || c.X >= 12 {
		return placeOK, gb
	}

	block := gb.Board[c.Y][c.X]
	log.Default().Printf("PreviewPlaceRock: %v at %s\n", block, c)
	if block.Marker == OpenBlock.Marker {
		gb.Board[c.Y][c.X].HighlightOK = true
		placeOK = true
	} else {
		gb.Board[c.Y][c.X].HighlightErr = true
	}
	gb.Board[c.Y][c.X].Marker = "X"

	return placeOK, gb
}

func (gb *Gameboard) GrabPiece() *LandPiece {
	grabPos := Coordinate{0, 0}
	grabPieceIndex, ngb := gb.PreviewGrabPiece(gb.CurrentPlayer(), grabPos)

	for {
		selectMenu := GetMovePieceMenu("Select piece")
		selectMenu.Options = append(selectMenu.Options, Option{Display: "Grab [N]one", ActionKey: "N"})
		if grabPieceIndex >= 0 {
			selectMenu.Options = append(selectMenu.Options, Option{Display: "[G]rab piece", ActionKey: "G"})
		}

		ngb.Display()
		pieceMove := GetPlayerMove(*gb.CurrentPlayer(), selectMenu)
		if IsValidMove(pieceMove, selectMenu) {
			switch strings.ToUpper(pieceMove) {
			case "N": // Select NONE
				log.Default().Printf("No land piece grabbed")
				return nil
			case "G": // Grab
				log.Default().Printf("Land piece %s grabbed from %s\n", (*gb.LandPieces)[grabPieceIndex], grabPos)
				// Remove Player's Grab card
				gb.CurrentPlayer().GrabCards = gb.CurrentPlayer().GrabCards[1:]
				gb.UpdateCards()
				gb.RemoveLandPieceFromGameboard(grabPieceIndex)
				gb.SetNextLandPieceTo(grabPieceIndex)
				newLandPiece := &(*gb.LandPieces)[grabPieceIndex]
				return newLandPiece
			case "R": // Right
				if (grabPos.X + 1) < 12 {
					grabPos.X++
				} else {
					selectMenu.RemoveOption("R")
				}
				grabPieceIndex, ngb = gb.PreviewGrabPiece(gb.CurrentPlayer(), grabPos)
			case "L": // Left
				if (grabPos.X - 1) >= 0 {
					grabPos.X--
				} else {
					selectMenu.RemoveOption("L")
				}
				grabPieceIndex, ngb = gb.PreviewGrabPiece(gb.CurrentPlayer(), grabPos)
			case "U": // Up
				if (grabPos.Y - 1) >= 0 {
					grabPos.Y--
				} else {
					selectMenu.RemoveOption("D")
				}
				grabPieceIndex, ngb = gb.PreviewGrabPiece(gb.CurrentPlayer(), grabPos)
			case "D": // Down
				if (grabPos.Y + 1) < 12 {
					grabPos.Y++
				} else {
					selectMenu.RemoveOption("D")
				}
				grabPieceIndex, ngb = gb.PreviewGrabPiece(gb.CurrentPlayer(), grabPos)
			}
		}
	}
}

// Place marker on board, highlight piece if it's up for grabs,
// return index of LandPiece if grab-able, updated gameboard
func (gb Gameboard) PreviewGrabPiece(p *Player, c Coordinate) (int, Gameboard) {
	grabPieceValue := uint16(0)
	grabPieceIdx := -1

	// cannot select at coordinate
	if c.Y < 0 || c.Y >= 12 || c.X < 0 || c.X >= 12 {
		return 0, gb
	}

	// TODO: grab replaces placed Rock 'R' with '#'
	for y := c.Y; y < 12 && grabPieceValue <= 0; y++ {
		for x := c.X; x < 12 && grabPieceValue <= 0; x++ {
			if gb.Board[c.Y][c.X].LandPieceValue > 0 {
				if grabPieceValue == 0 {
					grabPieceValue = gb.Board[c.Y][c.X].LandPieceValue
				}
			}
		}
	}

	block := gb.Board[c.Y][c.X]
	log.Default().Printf("PreviewGrabPiece: %v at %s\n", block, c)
	if block.Marker == LandPieceBlock.Marker && (block.Belongs_to.Name != p.Name) {
		gb.Board[c.Y][c.X].HighlightOK = true
		for idx, lp := range *gb.LandPieces {
			if lp.Value == grabPieceValue {
				log.Default().Printf("Found Land piece %d at index %d placed at:%s\n", grabPieceValue, idx, lp.PlacedAt)
				_, _, gb = gb.PlacePiecePreview(gb.CurrentPlayer(), lp, *lp.PlacedAt, true)
				grabPieceIdx = idx
			}
		}
	}
	gb.Board[c.Y][c.X].Marker = "X"

	return grabPieceIdx, gb
}

func (gb *Gameboard) SwapPiece() {
	selectedIdx := 0
	for selectedIdx <= 0 {
		swapMenu := Menu{Category: "Select piece", Options: []Option{}}

		for i := 1; i <= gb.showNumPieces; i++ {
			swapMenu.Options = append(swapMenu.Options, Option{Display: fmt.Sprintf("[%d]", i), ActionKey: fmt.Sprintf("%d", i)})
		}

		gb.Display()
		pieceMove := GetPlayerMove(*gb.CurrentPlayer(), swapMenu)
		if IsValidMove(pieceMove, swapMenu) {
			newIdx, err := strconv.Atoi(pieceMove)
			if err == nil {
				selectedIdx = newIdx
				log.Default().Printf("Current piece: %d", gb.currentPieceIndex)
				gb.currentPieceIndex = gb.currentPieceIndex + newIdx - 1
				log.Default().Printf("New piece selected: %d", gb.currentPieceIndex)
				gb.CurrentPlayer().SwapCards = gb.CurrentPlayer().SwapCards[1:]
				gb.UpdateCards()
			}
		}
	}
}

func (gb *Gameboard) PlacePiece(curLandPiece *LandPiece) *Gameboard {
	placeXY := Coordinate{X: 0, Y: 0}
	keepPlacing := true
	tmpP := Player{Name: gb.CurrentPlayer().Name, Color: Green}
	itFits, valid, ngb := gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
	var moveMenu Menu

	for keepPlacing {
		ngb.Display()
		if itFits {
			moveMenu = GetMovePieceMenu("Move piece")
		}
		if valid {
			moveMenu.Options = append(moveMenu.Options, Option{Display: "[P]lace piece", ActionKey: "P"})
		}
		rotateMenu := Menu{Category: "Rotate piece", Options: []Option{
			{Display: "[C]lockwise", ActionKey: "C"},
			{Display: "[A]nti-clockwise", ActionKey: "A"},
		}}
		pieceMove := GetPlayerMove(*gb.CurrentPlayer(), moveMenu, rotateMenu)
		if IsValidMove(pieceMove, moveMenu, rotateMenu) {
			var previewBoard Gameboard
			switch strings.ToUpper(pieceMove) {
			case "P": // Place
				curLandPiece.PlacedAt = &placeXY
				ngb.currentPieceIndex++
				log.Default().Printf("Place land piece: mem:%p, placed:%v\n", curLandPiece, curLandPiece.PlacedAt)
				keepPlacing = false
				ngb.ReplacePlayer(&tmpP, gb.CurrentPlayer())
				ngb.UpdateCards()
				return &ngb
			case "R": // Right
				placeXY.X++
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				if !itFits {
					placeXY.X--
					moveMenu.RemoveOption("R")
				}
			case "L": // Left
				placeXY.X--
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				if !itFits {
					placeXY.X++
					moveMenu.RemoveOption("L")
				}
			case "U": // Up
				placeXY.Y--
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				if !itFits {
					placeXY.Y++
					moveMenu.RemoveOption("U")
				}
			case "D": // Down
				placeXY.Y++
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				if !itFits {
					placeXY.Y--
					moveMenu.RemoveOption("D")
				}
			case "C": // rotate Clockwise
				curLandPiece.Value, _ = RotateClockwise(curLandPiece.Value)
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				//TODO: can not currently happen, undo rotate when it does
				// if !itFits {
				// }
			case "A": // rotate AntiClockwise
				curLandPiece.Value, _ = RotateAntiClockwise(curLandPiece.Value)
				itFits, valid, previewBoard = gb.PlacePiecePreview(&tmpP, *curLandPiece, placeXY, false)
				//TODO: can not currently happen, undo rotate when it does
				// if !itFits {
				// }
			}
			if itFits {
				ngb = previewBoard
			}
		} else {
			log.Default().Printf("Invalid option: %s\n", pieceMove)
		}
	}

	return gb
}

// PlacePiecePreview checks if a LandPiece fits on Gameboard at provided Coordinate
// Then renders the Gameboard depicting the new LandPiece at the provided Coordinate
// Returns confirmation if the piece could fit, if board in in valid state & the new Gameboard
func (gb Gameboard) PlacePiecePreview(p *Player, lp LandPiece, c Coordinate, overlapOK bool) (bool, bool, Gameboard) {
	lpStr := lp.String()
	boardValid := true

	// TODO: replace below with g.LandPieceFits() & cater for moving outside GameBoard for smaller pieces
	if c.Y < 0 || c.Y > 8 || c.X < 0 || c.X > 8 {
		log.Default().Printf("Coord %s does not fit\n", c)
		return false, false, gb
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			sidx := x + (y * 4)
			if lpStr[sidx:sidx+1] == "#" && (c.Y+y >= 0) && (c.X+x >= 0) && (c.Y+y < 16) && (c.X+x < 16) {
				block := gb.Board[c.Y+y][c.X+x]
				marker := block.Marker
				if block.Belongs_to == nil || (block.Belongs_to.Name == p.Name) {
					if strings.Trim(marker, " ") == OpenBlock.Marker {
						marker = LandPieceBlock.Marker
					}
					gb.Board[c.Y+y][c.X+x] = Block{Marker: marker, Belongs_to: p, LandPieceValue: lp.Value}
				} else {
					if overlapOK {
						gb.Board[c.Y+y][c.X+x] = Block{Marker: marker, HighlightOK: true}
					} else {
						gb.Board[c.Y+y][c.X+x] = Block{Marker: marker, HighlightErr: true}
					}
					boardValid = false
				}
			}
		}
	}

	// fmt.Println(g)

	return true, boardValid, gb // it fits
}

func GetCardMenu(p Player) Menu {
	cardMenu := Menu{Category: "Use a card", Options: []Option{}}
	if len(p.GrabCards) > 0 {
		cardMenu.Options = append(cardMenu.Options, Option{Display: "[G]rab land", ActionKey: "G"})
	}
	if len(p.SwapCards) > 0 {
		cardMenu.Options = append(cardMenu.Options, Option{Display: "[S]wap piece", ActionKey: "S"})
	}
	if len(p.RockCards) > 0 {
		cardMenu.Options = append(cardMenu.Options, Option{Display: "Place [R]ock", ActionKey: "R"})
	}

	return cardMenu
}

func (gb *Gameboard) UpdateCards() {
	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			block := gb.Board[y][x]
			if block.Belongs_to != nil && block.Marker != OpenBlock.Marker && block.Marker != LandPieceBlock.Marker && block.Marker != HomeBlock.Marker {
				typ := ""
				switch block.Marker {
				case SwapBlock.Marker:
					typ = "SWAP"
					block.Belongs_to.SwapCards = append(block.Belongs_to.SwapCards, SwapCard)
				case GrabBlock.Marker:
					typ = "GRAB"
					block.Belongs_to.GrabCards = append(block.Belongs_to.GrabCards, SwapCard)
				case RockBlock.Marker:
					typ = "ROCK"
					block.Belongs_to.RockCards = append(block.Belongs_to.RockCards, RockCard)
				}
				log.Default().Printf("Player %s gains a %s card\n", block.Belongs_to.Name, typ)
				gb.Board[y][x].Marker = LandPieceBlock.Marker
			}
		}
	}
}

func (gb Gameboard) GetRandomPosition(area Area) (int, int) {
	for {
		x := rand.Intn(area.End.X-area.Start.X) + area.Start.X
		y := rand.Intn(area.End.Y-area.Start.Y) + area.Start.Y

		if gb.Board[y][x].Marker == OpenBlock.Marker {
			return x, y
		} else {
			log.Default().Printf("%d,%d occupied by %s", x, y, gb.Board[y][x].Belongs_to)
		}
	}
}

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

type Area struct {
	Start Coordinate
	End   Coordinate
}

func (a Area) String() string {
	return fmt.Sprintf("%s->%s", a.Start, a.End)
}

func (a Area) Size() int {
	return ((a.End.X - a.Start.X) + 1) * ((a.End.Y - a.Start.Y) + 1)
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

func GetMovePieceMenu(category string) Menu {
	return Menu{Category: category, Options: []Option{
		{Display: "[L]eft", ActionKey: "L"},
		{Display: "[R]ight", ActionKey: "R"},
		{Display: "[U]p", ActionKey: "U"},
		{Display: "[D]own", ActionKey: "D"},
	}}
}
