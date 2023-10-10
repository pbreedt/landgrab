package landgame

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
	result := RotateClockwise(1234)
	expect := ".#.. | \n.##. | \n#... | \n.#.. | \n"
	t.Logf("\n%s", result)
	if result != expect {
		t.Errorf("expected:\n%s, got:\n%s", expect, result)
	}
}

func TestRotateAntiClockwise(t *testing.T) {
	result := RotateAntiClockwise(1234)
	expect := "..#. | \n...# | \n.##. | \n..#. | \n"
	t.Logf("\n%s", result)
	if result != expect {
		t.Errorf("expected:\n%s, got:\n%s", expect, result)
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
