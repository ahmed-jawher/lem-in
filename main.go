package main

import (
	"fmt"
	errordata "lem-in/errordata"
	pathsteps "lem-in/pathsteps"
	"os"
)

func main() {
    Arg := os.Args
    if len(Arg) != 2 {
        fmt.Println("Usage Example: go run . example00.txt")
    } else {
        errordata.ParseFile("Audit/" + Arg[1])
        pathsteps.FindAllPaths()
        pathsteps.FindCombinations()
        combinationh, AntsInPath := pathsteps.CheapestCombination()
        errordata.Moves(combinationh, AntsInPath)
    }

}
