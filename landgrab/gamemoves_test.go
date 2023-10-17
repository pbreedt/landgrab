package landgrab

import (
	"testing"
)

func TestSinglMenuDisplay(t *testing.T) {
	menu := Menu{Category: "Category", Options: []Option{{Display: "Option 1", ActionKey: "1"}, {Display: "Option 2", ActionKey: "2"}}}
	expect := "Category: Option 1 | Option 2\n"
	got := makeMenuDisplay(menu)
	if got != expect {
		t.Errorf("expected:\n'%s', got:\n'%s'", expect, got)
	}
}

func TestSingleNumMenuValid(t *testing.T) {
	menu := Menu{Category: "Category", Options: []Option{{Display: "Option 1", ActionKey: "1"}, {Display: "Option 2", ActionKey: "2"}}}
	expect := true
	got := IsValidMove("1", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
	got = IsValidMove("2", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}

	expect = false
	got = IsValidMove("3", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
}
func TestSingleAlphaMenuValid(t *testing.T) {
	menu := Menu{Category: "Category", Options: []Option{{Display: "Option A", ActionKey: "A"}, {Display: "Option B", ActionKey: "B"}}}
	expect := true
	got := IsValidMove("A", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
	got = IsValidMove("b", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}

	expect = false
	got = IsValidMove("C", menu)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
}

func TestMultiMenuDisplay(t *testing.T) {
	menu1 := Menu{Category: "Category1", Options: []Option{{Display: "Option 1.1", ActionKey: "1"}, {Display: "Option 1.2", ActionKey: "2"}}}
	menu2 := Menu{Category: "Category2", Options: []Option{{Display: "Option 2.1", ActionKey: "1"}, {Display: "Option 2.2", ActionKey: "2"}}}
	expect := "Category1: Option 1.1 | Option 1.2\nCategory2: Option 2.1 | Option 2.2\n"
	got := makeMenuDisplay(menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%s', got:\n'%s'", expect, got)
	}
}

func TestMultiMenuValid(t *testing.T) {
	menu1 := Menu{Category: "Category1", Options: []Option{{Display: "Option 1.1", ActionKey: "11"}, {Display: "Option 1.2", ActionKey: "12"}}}
	menu2 := Menu{Category: "Category2", Options: []Option{{Display: "Option 2.1", ActionKey: "21"}, {Display: "Option 2.2", ActionKey: "22"}}}
	expect := true
	got := IsValidMove("11", menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
	got = IsValidMove("12", menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
	got = IsValidMove("21", menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
	got = IsValidMove("22", menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}

	expect = false
	got = IsValidMove("3", menu1, menu2)
	if got != expect {
		t.Errorf("expected:\n'%t', got:\n'%t'", expect, got)
	}
}
