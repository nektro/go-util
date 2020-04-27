package cli

import (
	"errors"
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/nektro/go-util/ansi"
)

// Radio allows you to prompt the user for their choice given a list of options.
// Control keys are Up and Down arrows, and Enter.
// Esc exits early and throws error.
func Radio(options []string) (string, error) {
	if len(options) == 0 {
		return "", nil
	}
	err := keyboard.Open()
	if err != nil {
		return "", nil
	}
	defer keyboard.Close()
	selected := 0
	for {
		for i, item := range options {
			if i == selected {
				fmt.Print("[x] ")
			} else {
				fmt.Print("[ ] ")
			}
			fmt.Println(item)
		}
		_, key, err := keyboard.GetKey()
		if err != nil {
			return "", nil
		}
		if key == keyboard.KeyArrowDown {
			selected++
		}
		if key == keyboard.KeyArrowUp {
			selected--
		}
		if selected == len(options) {
			selected = 0
		}
		if selected == -1 {
			selected = len(options) - 1
		}
		for range options {
			fmt.Print(ansi.CursorUp(1))
			fmt.Print(ansi.EraseInLine(0))
		}
		if key == keyboard.KeyEsc {
			return "", errors.New("early exit")
		}
		if key == keyboard.KeyEnter {
			break
		}
	}
	return options[selected], nil
}
