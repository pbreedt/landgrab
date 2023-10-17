package landgrab

import (
	"fmt"
	"strings"

	"github.com/pbreedt/stdio/input"
)

type Menu struct {
	Category string
	Options  []Option
}

type Option struct {
	Display   string
	ActionKey string
}

func GetPlayerMove(p Player, s ...Menu) string {
	menuDisplay := makeMenuDisplay(s...)

	move, _ := input.ReadString(menuDisplay + fmt.Sprintf("Move for player %s? ", p.Name))

	return move
}
func makeMenuDisplay(s ...Menu) string {
	menuDisplay := ""

	for _, m := range s {
		options := collectDisplayValues(m)
		menuDisplay += fmt.Sprintf("%-15s: ", m.Category) + strings.Join(options, " | ") + "\n"
	}

	return menuDisplay
}

func collectDisplayValues(m Menu) []string {
	var values []string
	for _, o := range m.Options {
		values = append(values, o.Display)
	}

	return values
}

func IsValidMove(move string, s ...Menu) bool {
	for _, m := range s {
		if m.containsKey(move) {
			return true
		}
	}
	return false
}

func (m Menu) containsKey(k string) bool {
	for _, o := range m.Options {
		if strings.ToUpper(o.ActionKey) == strings.ToUpper(k) {
			return true
		}
	}

	return false
}

func (m Menu) RemoveOption(actionKey string) Menu {
	new := Menu{Category: m.Category}
	for _, o := range m.Options {
		if o.ActionKey != actionKey {
			new.Options = append(new.Options, o)
		}
	}

	return new
}
