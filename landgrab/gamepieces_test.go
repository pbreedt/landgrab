package landgrab

import (
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
		t.Errorf("expect no error, got %v", err)
	}
	if result != 9 {
		t.Errorf("expected:9, got:%d", result)
	}

	result, err = BinaryStringToInt("1111111111111111")
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if result != 65535 {
		t.Errorf("expected:65535, got:%d", result)
	}
}
