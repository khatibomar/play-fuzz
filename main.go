package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type matchingMode int8

const (
	same matchingMode = iota
	sensitive
	insensitive
)

var matchingAlgoMap map[string]matchingMode = map[string]matchingMode{
	"same":        same,
	"sensitive":   sensitive,
	"insensitive": insensitive,
}

var algoMap map[matchingMode]func(s1, s2 string) bool = map[matchingMode]func(s1 string, s2 string) bool{
	same: func(s1, s2 string) bool {
		return s1 == s2
	},
	sensitive: func(s1, s2 string) bool {
		return strings.Contains(s1, s2)
	},
	insensitive: func(s1, s2 string) bool {
		return strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
	},
}

type application struct {
	query        string
	reader       io.Reader
	pipeMode     bool
	matchingMode matchingMode
}

func getMatches(lines []string, query string, algo matchingMode) []string {
	result := make([]string, 0, len(query))
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
	flag.StringVar(&app.query, "query", "", "the sub string we need to match on [case insensitive]")
	flag.StringVar(&algo, "algo", "insensitive", "one of [same,sensitive,insensitive]")
	flag.Parse()

	var ok bool
	if app.matchingMode, ok = matchingAlgoMap[algo]; !ok {
		panic(fmt.Errorf("invalid algo, it must be one of [same,sensitive,insensitive]"))
	}

	if app.query == "" {
		flag.Usage()
		return
	}

	stat, _ := os.Stdin.Stat()
	app.pipeMode = (stat.Mode() & os.ModeCharDevice) == 0

	if err := app.run(); err != nil {
		panic(err)
	}
}

func (app *application) getInput() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(app.reader)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (app *application) run() error {
	if !app.pipeMode {
		fmt.Printf("[to stop enter empty line]\nInput:\n")
	}

	input, err := app.getInput()
	if err != nil {
		return fmt.Errorf("failed to get input: %w", err)
	}

	fmt.Println("processing matches...")
	matches := getMatches(input, app.query, app.matchingMode)
	if len(matches) == 0 {
		fmt.Println("no matches :(")
	} else {
		fmt.Println("I found matches :)")
		for i, m := range matches {
			fmt.Printf("match %d: %s\n", i+1, m)
		}
	}

	return nil
}
