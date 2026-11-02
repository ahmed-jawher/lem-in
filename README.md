# Lem-in

A digital ant farm solver written in Go. Given a colony of rooms and tunnels, it finds the set of paths that moves every ant from `##start` to `##end` in the **fewest possible turns**, then prints the movement turn by turn.

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Algorithms](https://img.shields.io/badge/topic-graph%20algorithms-blueviolet)

## Example

**Input** (`Audit/example00.txt`)

```text
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

**Output** (the input is echoed first, then one line per turn)

```text
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```

`L<ant>-<room>` means ant `<ant>` moved into room `<room>` during that turn. Each room holds only one ant at a time, except for the start and end rooms.

## Features

- **Parser and validator**: reads the ant count, rooms (`name x y`), tunnels (`a-b`) and the `##start` / `##end` commands
- **Strict input checks**: exactly one start and one end, no duplicate rooms or tunnels, no self-links, no links to unknown rooms, and a path from start to end must exist
- **Path search**: finds every route from start to end with depth-first search
- **Path combination**: builds sets of non-overlapping paths that share no rooms except the end
- **Optimizer**: spreads the ants across each set of paths and keeps the set that needs the fewest turns
- **Simulator**: moves the ants turn by turn and prints the moves sorted by ant ID

## Getting Started

**Requirements:** Go 1.20 or newer

```bash
git clone https://github.com/ahmed-jawher/lem-in.git
cd lem-in
go run . example00.txt
```

Input files are read from the `Audit/` folder. To run every bundled example:

```bash
bash test.sh
```

Invalid input stops the program with a descriptive error:

```text
$ go run . badexample00.txt
Error: invalid data format, invalid number of ants
```

## Input Format

| Line | Meaning |
|---|---|
| `4` | Number of ants (the first line; must be positive) |
| `##start` / `##end` | Marks the next room as the start or the end |
| `name x y` | A room and its coordinates |
| `a-b` | A two-way tunnel between rooms `a` and `b` |
| `# comment` | Ignored |

## Project Structure

```text
lem-in/
├── main.go          # entry point: parse → find paths → choose best → simulate
├── core/            # shared data structures and error helpers
├── errordata/       # input parsing, validation and move printing
├── pathsteps/       # path search, combinations and turn calculation
├── Audit/           # example and invalid input files
└── test.sh          # runs all examples
```

## Author

**Ahmed Jawher** — [@ahmed-jawher](https://github.com/ahmed-jawher)
