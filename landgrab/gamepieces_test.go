package landgrab

import (
	"strings"
	"testing"
)

func TestToBinaryGrid(t *testing.T) {
	result := ToBinaryGrid(1234)
	expect := ".... | \n.#.. | \n##.# | \n..#. | \n"
	t.Logf("\n%s", result)
	if result != expect {
		t.Errorf("expected:\n'%s', got:\n'%s'", expect, result)
	}
}

func TestRotateClockwise(t *testing.T) {
	nval, result := RotateClockwise(1234)
	expectRes := ".#.. | \n.##. | \n#... | \n.#.. | \n"
	expectVal := uint16(18052)
	t.Logf("\n%s", result)
	if result != expectRes {
		t.Errorf("expected:\n%s, got:\n%s", expectRes, result)
	}

	if nval != expectVal {
		t.Errorf("expected:\n%d, got:\n%d", expectVal, nval)
	}
}

func TestRotateAntiClockwise(t *testing.T) {
	nval, result := RotateAntiClockwise(1234)
	expectRes := "..#. | \n...# | \n.##. | \n..#. | \n"
	expectVal := uint16(8546)
	t.Logf("\n%s", result)
	if result != expectRes {
		t.Errorf("expected:\n%s, got:\n%s", expectRes, result)
	}

	if nval != expectVal {
		t.Errorf("expected:\n%d, got:\n%d", expectVal, nval)
	}
}

func TestBinaryString(t *testing.T) {
	result, err := BinaryStringToInt("1001")
	if err != nil {
		t.Errorf("expect no error, got error %v", err)
	}
	if result != 9 {
		t.Errorf("expected:9, got:%d", result)
	}

	result, err = BinaryStringToInt("1111111111111111")
	if err != nil {
		t.Errorf("expect no error, got error %v", err)
	}
	if result != 65535 {
		t.Errorf("expected:65535, got:%d", result)
	}
}

func doLandPieceEmptyTest(t *testing.T, test string, testData string) (bool, uint16) {
	result, err := BinaryStringToInt(strings.ReplaceAll(testData, " ", ""))
	if err != nil {
		t.Errorf("expect no error, got error %v", err)
	}
	landPiece := LandPiece{Value: result}
	switch test {
	case "LC":
		return landPiece.LeftColEmpty(), landPiece.Value
	case "RC":
		return landPiece.RightColEmpty(), landPiece.Value
	case "TR":
		return landPiece.TopRowEmpty(), landPiece.Value
	case "BR":
		return landPiece.BottomRowEmpty(), landPiece.Value
	}
	return false, 0
}

func TestLandPieceColEmpty(t *testing.T) {
	// expect LeftColEmpty TRUE
	testData := []string{
		"0111 0111 0111 0111",
		"0000 0000 0000 0000",
	}

	for _, td := range testData {
		leftColEmpty, value := doLandPieceEmptyTest(t, "LC", td)
		if !leftColEmpty {
			t.Errorf("expected: LeftColEmpty TRUE, got: FALSE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect LeftColEmpty FALSE
	testData = []string{
		"1111 1111 1111 1111",
		"1000 0000 0000 0000",
		"0000 1000 0000 0000",
		"0000 0000 1000 0000",
		"0000 0000 0000 1000",
		"1000 1000 1000 1000",
	}
	for _, td := range testData {
		leftColEmpty, value := doLandPieceEmptyTest(t, "LC", td)
		if leftColEmpty {
			t.Errorf("expected: LeftColEmpty FALSE, got: TRUE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect RightColEmpty TRUE
	testData = []string{
		"1110 1110 1110 1110",
		"0000 0000 0000 0000",
	}
	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "RC", td)
		if !empty {
			t.Errorf("expected: RightColEmpty TRUE, got: FALSE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect RightColEmpty FALSE
	testData = []string{
		"1111 1111 1111 1111",
		"0000 0000 0000 0001",
		"0000 0000 0001 0000",
		"0000 0001 0000 0000",
		"0001 0000 0000 0000",
		"0001 0001 0001 0001",
	}
	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "RC", td)
		if empty {
			t.Errorf("expected: RightColEmpty FALSE, got: TRUE for piece: %s", ToBinaryGrid(value))
		}
	}
}

func TestLandPieceRowEmpty(t *testing.T) {
	// expect TopRowEmpty TRUE
	testData := []string{
		"0000 1111 1111 1111",
		"0000 0000 0000 0000",
	}

	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "TR", td)
		if !empty {
			t.Errorf("expected: TopRowEmpty TRUE, got: FALSE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect TopRowEmpty FALSE
	testData = []string{
		"1111 1111 1111 1111",
		"1000 0000 0000 0000",
		"0100 0000 0000 0000",
		"0010 0000 0000 0000",
		"0001 0000 0000 0000",
	}
	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "TR", td)
		if empty {
			t.Errorf("expected: TopRowEmpty FALSE, got: TRUE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect BottomRowEmpty TRUE
	testData = []string{
		"1111 1111 1111 0000",
		"0000 0000 0000 0000",
	}
	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "BR", td)
		if !empty {
			t.Errorf("expected: BottomRowEmpty TRUE, got: FALSE for piece: %s", ToBinaryGrid(value))
		}
	}

	// expect BottomRowEmpty FALSE
	testData = []string{
		"1111 1111 1111 1111",
		"0000 0000 0000 0001",
		"0000 0000 0000 0010",
		"0000 0000 0000 0100",
		"0000 0000 0000 1000",
	}
	for _, td := range testData {
		empty, value := doLandPieceEmptyTest(t, "BR", td)
		if empty {
			t.Errorf("expected: BottomRowEmpty FALSE, got: TRUE for piece: %s", ToBinaryGrid(value))
		}
	}
}
