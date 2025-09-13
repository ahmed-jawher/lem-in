package AntColony

import (
	"lem-in/core"
	"sort"
	"strings"
)

func FindAllPaths() {
    visited := make(map[string]bool)
    findPaths(core.ColonyInstance.RoomTunnels, core.ColonyInstance.StartRoom, core.ColonyInstance.EndRoom, []string{}, visited, core.PossiblePaths)
    if len(core.PossiblePaths) == 0 {
        core.HandleDataError("no path between start and end")
    }
}

func findPaths(graph map[string][]string, currentNode string, destination string, path []string, visited map[string]bool, paths [][]string) {
    path = append(path, currentNode)
    visited[currentNode] = true

	if currentNode == destination {
		// Append the completed path to the paths slice
		core.PossiblePaths = append(core.PossiblePaths, append([]string(nil), path...))
	} else {
		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				findPaths(graph, neighbor, destination, path, visited, paths)
			}
		}
	}

    path = path[:len(path)-1]
    visited[currentNode] = false
}
func FindCombinations() {
    for _, Path := range core.PossiblePaths {
        combinedPath := ComparePath(Path[1:])
        core.Combinations = append(core.Combinations, combinedPath)
    }

    FilterSimilarCombinations()
    core.Combinations = SortSubpathsByLength(core.Combinations)
}

func ComparePath(Path []string) core.CombinedPath {
	combinedPath := core.CombinedPath{
		Path: make([][]string, 0),
	}
	combinedPath.Path = append(combinedPath.Path, Path)

	var UsedNodes = Path
	for _, Path2 := range core.PossiblePaths {
		Path2 := Path2[1:]
		if !equalSlices(Path, Path2) {
			Used := false
			for _, Path2Nodes := range Path2 {
				for _, Nodes := range UsedNodes {
					if Path2Nodes == Nodes && Nodes != core.ColonyInstance.EndRoom {
						Used = true
						break
					}
				}
			}
			if !Used {
				combinedPath.Path = append(combinedPath.Path, Path2)
				UsedNodes = append(UsedNodes, Path2...)
			}
		}
	}

	return combinedPath
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FilterSimilarCombinations() {
    filteredCombinations := make([]core.CombinedPath, 0)
    visitedPaths := make(map[string]bool)

	for _, combination := range core.Combinations {
		sortedPath := getSortedPath(combination.Path)
		if !visitedPaths[sortedPath] {
			visitedPaths[sortedPath] = true
			filteredCombinations = append(filteredCombinations, combination)
		}
	}
	core.Combinations = filteredCombinations
}

func getSortedPath(path [][]string) string {
	sortedPath := make([]string, len(path))
	for i, subpath := range path {
		sortedSubpath := make([]string, len(subpath))
		copy(sortedSubpath, subpath)
		sort.Strings(sortedSubpath)
		sortedPath[i] = strings.Join(sortedSubpath, "-")
	}
	sort.Strings(sortedPath)
	return strings.Join(sortedPath, "#")
}

func SortSubpathsByLength(combinations []core.CombinedPath) []core.CombinedPath {
    for i := range combinations {
        sort.SliceStable(combinations[i].Path, func(j, k int) bool {
            return len(combinations[i].Path[j]) < len(combinations[i].Path[k])
        })
    }

    return combinations
}
