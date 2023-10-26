package landgrab

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type LandPiecesSlice []LandPiece

type LandPiece struct {
	// 0000
	// 0010
	// 0110
	// 0100
	Value    uint16 // 0000_0010_0110_0100
	PlacedAt *Coordinate
}

func RandomizeLandPieces(num int) *LandPiecesSlice {
	var lps LandPiecesSlice
	for i := 0; i < num; i++ {
		n := uint16(rand.Intn(math.MaxUint16))
		if !lps.Contains(n) {
			lps = append(lps, LandPiece{Value: n})
		}
	}
	return &lps
}

func (s LandPiecesSlice) Contains(num uint16) bool {
	for _, lp := range s {
		if lp.Value == num {
			return true
		}
	}
	return false
}

func (lps LandPiecesSlice) PrintUnplacedN(startIdx, n int) int {
	var nums []uint16
	firstPieceIdx := -1
	selectIdx := 1
	selectHdr := ""
	for i := startIdx; len(nums) < n && i < len(lps); i++ {
		// log.Default().Printf("Land piece:%s mem:%p, index=%d, placed:%v\n", lps[i], &lps[i], i, lps[i].PlacedAt)
		if lps[i].PlacedAt == nil {
			if firstPieceIdx < 0 {
				firstPieceIdx = i
			}
			nums = append(nums, lps[i].Value)
			selectHdr += fmt.Sprintf(" [%2d] |", selectIdx)
			selectIdx++
		}
	}

	fmt.Println(selectHdr[1:])
	fmt.Println(ToBinaryGrid(nums...))

	return firstPieceIdx
}

func (lp LandPiece) String() string {
	s := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", lp.Value), "0", "."), "1", "#")

	return s
}

func (lp LandPiece) TopRowEmpty() bool {
	binaryStr := fmt.Sprintf("%016b", lp.Value)

	return binaryStr[0:4] == "0000"
}
func (lp LandPiece) BottomRowEmpty() bool {
	binaryStr := fmt.Sprintf("%016b", lp.Value)

	return binaryStr[12:16] == "0000"
}
func (lp LandPiece) LeftColEmpty() bool {
	binaryStr := fmt.Sprintf("%016b", lp.Value)
	lc := binaryStr[0:1] + binaryStr[4:5] + binaryStr[8:9] + binaryStr[12:13]
	return lc == "0000"
}
func (lp LandPiece) RightColEmpty() bool {
	binaryStr := fmt.Sprintf("%016b", lp.Value)
	rc := binaryStr[3:4] + binaryStr[7:8] + binaryStr[11:12] + binaryStr[15:16]
	return rc == "0000"
}

func ToBinaryGrid(nums ...uint16) string {
	line1, line2, line3, line4 := "", "", "", ""
	for i, num := range nums {
		binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
		// fmt.Println("ToBinaryGrid before:", binaryStr)

		if i == 0 {
			line1 += ColorString(binaryStr[0:4], Red) + " | "
			line2 += ColorString(binaryStr[4:8], Red) + " | "
			line3 += ColorString(binaryStr[8:12], Red) + " | "
			line4 += ColorString(binaryStr[12:16], Red) + " | "
		} else {
			line1 += binaryStr[0:4] + " | "
			line2 += binaryStr[4:8] + " | "
			line3 += binaryStr[8:12] + " | "
			line4 += binaryStr[12:16] + " | "
		}

	}

	// Join the rows of the grid with spaces
	result := line1 + "\n" + line2 + "\n" + line3 + "\n" + line4 + "\n"

	return result
}

// 0000			1000
// 0010			1100
// 0110   ==>  	0110
// 1100			0000
func RotateClockwise(num uint16) (uint16, string) {
	binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
	binaryStrResult := ""

	// fmt.Println("RotateClockwise before:", binaryStr)

	grid := make([]string, 4)
	for i := 0; i < 4; i++ {
		// fmt.Printf("[%d]: [%d:%d] [%d:%d] [%d:%d] [%d:%d]\n", i, i+12, i+12+1, i+8, i+8+1, i+4, i+4+1, i, i+1)
		binaryStrResult += binaryStr[i+12:i+12+1] + binaryStr[i+8:i+8+1] + binaryStr[i+4:i+4+1] + binaryStr[i:i+1]
		grid[i] = binaryStr[i+12:i+12+1] + binaryStr[i+8:i+8+1] + binaryStr[i+4:i+4+1] + binaryStr[i:i+1] + " | "
	}

	// Join the rows of the grid with spaces
	result := strings.Join(grid, "\n") + "\n"
	// fmt.Println("RotateClockwise after:", result)

	binaryStrResult = strings.ReplaceAll(strings.ReplaceAll(binaryStrResult, ".", "0"), "#", "1")
	intVal, err := BinaryStringToInt(binaryStrResult)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	return intVal, result
}

// 0000			0000
// 0010			0110
// 0110   ==>  	0011
// 1100			0001
func RotateAntiClockwise(num uint16) (uint16, string) {
	binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
	binaryStrResult := ""

	// fmt.Println("RotateAntiClockwise before:", binaryStr)

	grid := make([]string, 4)
	for i := 3; i >= 0; i-- {
		// fmt.Printf("[%d]: [%d:%d] [%d:%d] [%d:%d] [%d:%d]\n", i, i+12, i+12+1, i+8, i+8+1, i+4, i+4+1, i, i+1)
		binaryStrResult += binaryStr[i:i+1] + binaryStr[i+4:i+4+1] + binaryStr[i+8:i+8+1] + binaryStr[i+12:i+12+1]
		grid[3-i] = binaryStr[i:i+1] + binaryStr[i+4:i+4+1] + binaryStr[i+8:i+8+1] + binaryStr[i+12:i+12+1] + " | "
	}

	// Join the rows of the grid with spaces
	result := strings.Join(grid, "\n") + "\n"
	// fmt.Println("RotateAntiClockwise after:", result)

	binaryStrResult = strings.ReplaceAll(strings.ReplaceAll(binaryStrResult, ".", "0"), "#", "1")
	intVal, err := BinaryStringToInt(binaryStrResult)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	return intVal, result
}

//TODO: implement FlipH & FlipV
//func FlipH(num uint16) string {}

func BinaryStringToInt(binaryStr string) (uint16, error) {
	if i, err := strconv.ParseInt(binaryStr, 2, 32); err != nil {
		return 0, err
	} else {
		return uint16(i), nil
	}
}

// func SingleToBinaryGrid(num uint16) string {
// 	// Convert the number to a binary string
// 	binaryStr := fmt.Sprintf("%016b", num)

// 	binaryStr = strings.ReplaceAll(binaryStr, "0", "_")
// 	binaryStr = strings.ReplaceAll(binaryStr, "1", "#")

// 	// Create a 4x4 grid from the binary string
// 	grid := make([]string, 4)
// 	for i := 0; i < 4; i++ {
// 		start := i * 4
// 		end := start + 4
// 		grid[i] = binaryStr[start:end]
// 	}

// 	// Join the rows of the grid with spaces
// 	result := strings.Join(grid, "\n")

//		return result
//	}
