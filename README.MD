Lem-in (Go)

Ant colony path routing solver for the Lem-in project, implemented in Go. It parses a graph of rooms and tunnels, finds all disjoint path combinations from start to end, selects the combination that minimizes the total number of turns for all ants, and prints the ants’ movements.

Features
- Parses Lem-in input: ants count, rooms, start/end markers, and tunnels.
- Validates input: single start/end, no self-links, no duplicate rooms/tunnels, links only between known rooms.
- Finds all paths and builds compatible path sets (no overlapping rooms except the end).
- Chooses the fastest combination using a simple load-balancing step calculator.
- Outputs ant moves per turn in the standard `L<id>-<room>` format.

Requirements
- Go 1.20+

Repository Layout
- `main.go`: CLI entrypoint and high-level flow.
- `core/`: shared structs and error helpers.
- `errordata/`: input parsing, validation, and movement output.
- `pathsteps/`: path search, combination generation, step calculation.
- `Audit/`: example input files used by the test script.
- `test.sh`: quick runner over multiple sample inputs.

Input Format
- First non-comment line: number of ants (positive integer).
- Rooms: `<name> <x> <y>` (coordinates are accepted but not used by the solver).
- Tunnels: `<room1>-<room2>`
- Directives: `##start` and `##end` placed immediately before the corresponding room line.
- Comments: lines starting with `#` (except `##start` and `##end`) are ignored. Blank lines are ignored.

Notes on validation
- Exactly one `##start` and one `##end` room must be present.
- All tunnels must reference known rooms.
- No room can link to itself.
- No duplicate room definitions or duplicate tunnels.
- At least one path from start to end must exist.

Build and Run
- From the repository root, inputs are resolved relative to `Audit/`.

- Run a single example
  - `go run . example00.txt`

- Build a binary
  - `go build -o lem-in`
  - `./lem-in example01.txt`

- Run all bundled samples
  - `bash test.sh`

Expected Output
1) The validated input is echoed (comments are skipped except for `##start` / `##end`).
2) Each subsequent line represents a turn, listing space-separated moves sorted by ant ID:
   - `L1-roomB L2-roomC ...`

Minimal Example
Given this input (saved under `Audit/example_min.txt`):

```
3
##start
A 0 0
B 1 0
##end
C 2 0
A-B
B-C
```

Run `go run . example_min.txt` and you will see the file echoed followed by turns similar to:

```
3
##start
A 0 0
B 1 0
##end
C 2 0
A-B
B-C
 L1-B
 L1-C L2-B
 L2-C L3-B
 L3-C
```

How It Works (Brief)
- Parse: builds a graph (`core.ColonyInstance`) and validates the dataset.
- Paths: DFS to enumerate all possible paths from start to end.
- Combine: generates sets of non-overlapping paths; filters duplicates; sorts paths shortest-first.
- Choose: distributes ants across a combination to minimize the total turns; selects the best combination.
- Move: simulates per-turn movement and prints moves in sorted order.

CLI Usage
- Usage: `go run . <input-file-in-Audit>`
- Example: `go run . example07.txt`

Troubleshooting
- Error: invalid data format, <reason>: The parser/validator found an issue in the input.
- If you pass a filename that is not inside `Audit/`, the program won’t find it (the code automatically prefixes `Audit/`).



