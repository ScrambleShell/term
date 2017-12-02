package main

import (
	"strings"

	"github.com/kiasaki/term"
)

func main() {
	t := term.NewTerminal()

	t.Start()
	defer t.Stop()

	prompt := "Hello! What's your name? "
	name := ""
	t.Puts(prompt)
	for {
		ev := <-t.Events()

		// Exit on Ctrl-C or Ctrl-D
		// Ctrl-C doesn't kill the program when the terminal is in raw mode
		if ev.Key == term.KeyCtrlC || ev.Key == term.KeyCtrlD {
			printOnNewLine(t, "Goodbye!", name)
			return

		}

		// Enter "submits" the entered name
		if ev.Key == term.KeyCr {
			printOnNewLine(t, "Nice to meet you %s!", name)
			return

		}

		// Remove the last character if possible and rewrite prompt line
		if ev.Key == term.KeyBackspace {
			if len(name) > 0 {
				name = name[:len(name)-1]
			}
			t.SetCursorColumn(0)
			t.Puts(strings.Repeat(" ", t.Width))
			t.SetCursorColumn(0)
			t.Puts(prompt + name)
			continue

		}

		// If the event is simply a character, append it to `name`
		if ev.Key == term.KeyRune {
			name += string(ev.Rune)
			t.Puts(string(ev.Rune))
		}
	}
}

func printOnNewLine(t *term.Terminal, f string, args ...interface{}) {
	t.Puts("\n")
	t.SetCursorColumn(0)
	t.Puts(f+"\n", args...)
	t.SetCursorColumn(0)
}
