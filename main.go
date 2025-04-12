package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

type matchingMode int8

const (
	same matchingMode = iota
	sensitive
	insensitive
	fuzzy
)

var matchingAlgoMap map[string]matchingMode = map[string]matchingMode{
	"same":        same,
	"sensitive":   sensitive,
	"insensitive": insensitive,
	"fuzzy":       fuzzy,
}

var algoMap map[matchingMode]func(text, query string) bool = map[matchingMode]func(s1 string, s2 string) bool{
	same: func(s1, s2 string) bool {
		return s1 == s2
	},
	sensitive: func(s1, s2 string) bool {
		return strings.Contains(s1, s2)
	},
	insensitive: func(s1, s2 string) bool {
		return strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
	},
	fuzzy: func(text, query string) bool {
		text = strings.ToLower(text)
		query = strings.ToLower(query)

		if len(query) == 0 {
			return true
		}
		if len(text) == 0 {
			return false
		}

		patternIdx := 0
		for i := 0; i < len(text) && patternIdx < len(query); i++ {
			if text[i] == query[patternIdx] {
				patternIdx++
			}
		}

		return patternIdx == len(query)
	},
}

type application struct {
	reader       io.Reader
	matchingMode matchingMode
}

func getMatches(lines []string, query string, algo matchingMode) []string {
	result := make([]string, 0, len(lines))
	for _, l := range lines {
		if algoMap[algo](l, query) {
			result = append(result, l)
		}
	}

	return result
}

func main() {
	app := &application{
		reader: os.Stdin,
	}

	var algo string
	flag.StringVar(&algo, "algo", "fuzzy", "one of [fuzzy,same,sensitive,insensitive]")
	flag.Parse()

	var ok bool
	if app.matchingMode, ok = matchingAlgoMap[algo]; !ok {
		fmt.Printf("invalid algo, it must be one of [same,sensitive,insensitive]\n")
		os.Exit(1)
	}

	if err := app.run(); err != nil {
		fmt.Printf("failed to run app: %v\n", err)
		os.Exit(1)
	}
}

func (app *application) run() error {
	var input []string

	stat, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat stdin: %w", err)
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return fmt.Errorf("no input provided (try piping something to stdin)")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}
	if len(input) == 0 {
		return fmt.Errorf("no input provided (empty pipe)")
	}

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("failed to open /dev/tty: %w", err)
	}
	defer tty.Close()

	oldState, err := term.MakeRaw(int(tty.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(tty.Fd()), oldState)

	screen := struct {
		io.Reader
		io.Writer
	}{tty, tty}
	t := term.NewTerminal(screen, "")

	var query strings.Builder
	prompt := string(t.Escape.Red) + "query> " + string(t.Escape.Reset)

	fmt.Fprint(t, "\x1b[H\x1b[K"+prompt+query.String())
	fmt.Fprint(t, "\x1b[2;1H\x1b[J")

	matches := getMatches(input, query.String(), app.matchingMode)

	if len(matches) > 0 {
		fmt.Fprint(t, string(t.Escape.Green))
		for _, match := range matches {
			fmt.Fprintf(t, "%s\n", match)
		}
		fmt.Fprint(t, string(t.Escape.Reset))
	}

	for {
		// Clear the entire query line before redrawing
		fmt.Fprint(t, "\x1b[H\x1b[K"+prompt+query.String())

		var b [1]byte
		if _, err := screen.Reader.Read(b[:]); err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		// Clear everything below the query line
		fmt.Fprint(t, "\x1b[2;1H\x1b[J")

		switch b[0] {
		case 3: // ctrl-c
			return nil
		case 13: // enter
			fmt.Fprintf(t, "\n%d matches:\n", len(matches))
			if len(matches) > 0 {
				fmt.Fprint(t, string(t.Escape.Green))
				for _, match := range matches {
					fmt.Fprintf(t, "%s\n", match)
				}
				fmt.Fprint(t, string(t.Escape.Reset))
			}
			return nil
		case 27: // escape
			return nil
		case 127: // backspace
			if query.Len() > 0 {
				str := query.String()
				query.Reset()
				query.WriteString(str[:len(str)-1])
				matches = getMatches(input, query.String(), app.matchingMode)
			}
		default:
			if b[0] >= 32 && b[0] < 127 { // printable characters only
				query.WriteByte(b[0])
				matches = getMatches(input, query.String(), app.matchingMode)
			}
		}

		if len(matches) > 0 {
			fmt.Fprint(t, string(t.Escape.Green))
			for _, match := range matches {
				fmt.Fprintf(t, "%s\n", match)
			}
			fmt.Fprint(t, string(t.Escape.Reset))
		}

		// Move cursor back to end of query
		fmt.Fprintf(t, "\x1b[1;%dH", len(prompt)+query.Len())
	}
}
