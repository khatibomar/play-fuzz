# play-fuzz

I am trying a new idea to have fun on weekends, I use AI to give me a tasks-list to implement something for fun.
The readme will be whatever the AI agent toss to me as "requirements"
and I implement them without using AI.

# 🧠 Terminal Fuzzy Finder – Phase 1

## 🎯 Goal
Build a basic CLI tool in Go that:
- Reads input lines from `stdin`
- Accepts a query from the command line
- Filters input lines by checking if they contain the query
- Prints matching results to `stdout`

---

## ✅ Tasks

### 📥 1. Read Input
- [X] Create a function to read all lines from stdin.
- [X] Use `bufio.Scanner` to read each line.
- [X] Store the lines in a slice (`[]string`).

### 🔍 2. Handle Query & Filtering
- [X] Get the search query from command-line arguments.
- [X] If no argument is provided, print usage instructions and exit.
- [X] Loop through all lines and check if each line contains the query (substring).
- [X] Make the search case-insensitive (convert both to lowercase before comparing).

### 📤 3. Output Results
- [X] Collect all matching lines into a new slice.
- [X] Print each matching line to `stdout`.

### 🧪 4. Manual Testing
- [X] Pipe input into the tool from another command (e.g. `ls | go run main.go foo`).
- [X] Confirm only matching lines are printed.
- [X] Test case sensitivity, empty queries, and no matches.

### 🧼 5. Optional Enhancements
- [X] Add a `--case-sensitive` flag.
- [X] Add support for exact matching (`==` instead of `contains`).
- [X] Add a `--help` flag or usage instructions.

---

