package landgame

import (
	"fmt"
	"strconv"
	"strings"
)

var LandPieces [100]LandPiece

type LandPiece struct {
	// 0000
	// 0010
	// 0110
	// 0100
	Bits    uint16 // 0000_0000_0000_0000
	placedX int
	placedY int
}

func (lp LandPiece) String() string {
	s := ""

	return s
}

func ToBinaryGrid(nums ...uint16) string {
	line1, line2, line3, line4 := "", "", "", ""
	for _, num := range nums {
		binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
		fmt.Println("ToBinaryGrid before:", binaryStr)

		line1 += binaryStr[0:4] + " | "
		line2 += binaryStr[4:8] + " | "
		line3 += binaryStr[8:12] + " | "
		line4 += binaryStr[12:16] + " | "
	}

	// Join the rows of the grid with spaces
	result := line1 + "\n" + line2 + "\n" + line3 + "\n" + line4 + "\n"

	return result
}

// 0000			1000
// 0010			1100
// 0110   ==>  	0110
// 1100			0000
func RotateClockwise(num uint16) string {
	binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
	fmt.Println("RotateClockwise before:", binaryStr)
	grid := make([]string, 4)
	for i := 0; i < 4; i++ {
		// fmt.Printf("[%d]: [%d:%d] [%d:%d] [%d:%d] [%d:%d]\n", i, i+12, i+12+1, i+8, i+8+1, i+4, i+4+1, i, i+1)
		grid[i] = binaryStr[i+12:i+12+1] + binaryStr[i+8:i+8+1] + binaryStr[i+4:i+4+1] + binaryStr[i:i+1] + " | "
	}

	// Join the rows of the grid with spaces
	result := strings.Join(grid, "\n") + "\n"

	return result
}

// 0000			0000
// 0010			0110
// 0110   ==>  	0011
// 1100			0001
func RotateAntiClockwise(num uint16) string {
	binaryStr := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%016b", num), "0", "."), "1", "#")
	fmt.Println("RotateClockwise before:", binaryStr)
	grid := make([]string, 4)
	for i := 3; i >= 0; i-- {
		// fmt.Printf("[%d]: [%d:%d] [%d:%d] [%d:%d] [%d:%d]\n", i, i+12, i+12+1, i+8, i+8+1, i+4, i+4+1, i, i+1)
		grid[3-i] = binaryStr[i:i+1] + binaryStr[i+4:i+4+1] + binaryStr[i+8:i+8+1] + binaryStr[i+12:i+12+1] + " | "
	}

	// Join the rows of the grid with spaces
	result := strings.Join(grid, "\n") + "\n"

	return result
}

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
