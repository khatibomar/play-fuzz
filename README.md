# play-fuzz

I am trying a new idea to have fun on weekends, I use AI to give me a tasks-list to implement something for fun.
The readme will be whatever the AI agent toss to me as "requirements"
and I implement them without using AI.

Feedback: chatgpt spits missing requirements so I had to correct it after it generates.
Like imagine telling it to do a fuzzy finder but it gave steps without the acutal fuzzy XD

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

# 🧠 Terminal Fuzzy Finder – Phase 2

## 🎯 Goal

Make the fuzzy finder interactive **and** actually fuzzy:

- Accept keyboard input in real-time (no Enter key needed)
- Filter results using a fuzzy matching algorithm
- Display the current query and matching lines
- Make everything update live as the user types

---

## ✅ Tasks

### 🧵 1. Enable Raw Terminal Mode

- [X] Use `golang.org/x/term` to enable raw input mode (no line buffering)
- [X] Save original terminal state and restore it on exit or interrupt
- [X] Clear and redraw the terminal every time the input changes

---

### 🎹 2. Capture Key Presses

- [X] Read single key presses (unbuffered)
- [X] Handle:
  - [X] Printable characters
  - [X] Backspace (`0x7f`)
  - [X] Escape (`\x1b`) or Ctrl+C to quit

---

### 🧠 3. Manage Internal State

- [X] Store:
  - Current query string
  - Input list (from stdin)
  - Filtered + scored matches

---

### 🧮 4. Implement Fuzzy Matching

- [X] Build a fuzzy scoring algorithm:
  - Match characters in order (non-contiguous allowed)
  - Higher score = closer match (shorter distance, earlier match, etc.)
  - Optionally use existing algo as a reference: `fzf`, `skim`, or `fzf-lite`
- [ ] Return a list of matched lines with scores
- [ ] Sort matches by descending score
- [ ] Limit result count (e.g., top 10–20)

---

### 🖥 5. Draw the UI

- [X] Clear screen using ANSI codes (`\033[H\033[2J`)
- [X] Draw:
  - Prompt line:  
    ```
    > your-query
    ```
  - List of best matches
- [X] Optionally:
  - Highlight matched characters in the results
  - Truncate long results with `...`

---

### 🧪 6. Manual Testing

- [X] Run:  
  ```bash
  cat biglist.txt | ./fuzzy
  ```
- [X] Type and confirm:
  - Matches appear in real-time
  - Fuzzy matching works for partial, non-contiguous queries
  - Handles fast typing, deletion, no matches
- [X] Exit using ESC or Ctrl+C

---

