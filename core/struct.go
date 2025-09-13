package core

var ColonyInstance Colony

type Colony struct {
	RoomTunnels map[string][]string
	StartRoom   string
	EndRoom     string
}

var PossiblePaths [][]string

type CombinedPath struct {
	Path [][]string
}

var Combinations []CombinedPath

var AntsNum int
