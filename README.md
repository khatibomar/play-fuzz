# play-fuzz

I am trying a new idea to have fun on weekends, I use AI to give me a tasks-list to implement something for fun.
The readme will be whatever the AI agent toss to me as "requirements"

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
- [ ] Create a function to read all lines from stdin.
- [ ] Use `bufio.Scanner` to read each line.
- [ ] Store the lines in a slice (`[]string`).

### 🔍 2. Handle Query & Filtering
- [ ] Get the search query from command-line arguments.
- [ ] If no argument is provided, print usage instructions and exit.
- [ ] Loop through all lines and check if each line contains the query (substring).
- [ ] Make the search case-insensitive (convert both to lowercase before comparing).

### 📤 3. Output Results
- [ ] Collect all matching lines into a new slice.
- [ ] Print each matching line to `stdout`.

### 🧪 4. Manual Testing
- [ ] Pipe input into the tool from another command (e.g. `ls | go run main.go foo`).
- [ ] Confirm only matching lines are printed.
- [ ] Test case sensitivity, empty queries, and no matches.

### 🧼 5. Optional Enhancements
- [ ] Add a `--case-sensitive` flag.
- [ ] Add support for exact matching (`==` instead of `contains`).
- [ ] Add a `--help` flag or usage instructions.

---

## 💡 Example Usage

```bash
# With files
cat file_list.txt | go run main.go term

# With command output
ls | go run main.go go

