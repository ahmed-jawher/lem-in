package AntColony

import (
	"bufio"
	"fmt"
	"lem-in/core"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) {
    core.ColonyInstance = core.Colony{
        RoomTunnels: make(map[string][]string),
        StartRoom:   "",
        EndRoom:     "",
    }
	NumOfStartRoom := 0
	NumOfEndroom := 0

	file, err := os.Open(filename)
	if err != nil {
		core.HandleError(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true

    for scanner.Scan() {
        line := scanner.Text()
        var Link []string
        var RoomName []string

        if firstLine {
            numAnts, err := strconv.ParseInt(line, 10, 64)
            if err != nil || numAnts <= 0 || numAnts>999999 {
                core.HandleDataError("invalid number of ants")
            }
            core.AntsNum = int(numAnts)
            firstLine = false
        } else if strings.HasPrefix(line, "##") {
            word := line[2:]
            var NextLine string
            var NextLineArray []string

            if scanner.Scan() {
                NextLine = scanner.Text()
                NextLineArray = strings.Split(NextLine, " ")
            }
            if word == "start" {
                core.ColonyInstance.StartRoom = NextLineArray[0]
                core.ColonyInstance.RoomTunnels[core.ColonyInstance.StartRoom] = make([]string, 0)
                NumOfStartRoom++
            } else if word == "end" {
                core.ColonyInstance.EndRoom = NextLineArray[0]
                core.ColonyInstance.RoomTunnels[core.ColonyInstance.EndRoom] = make([]string, 0)
                NumOfEndroom++
            }
        } else if strings.Contains(line, "-") {
            Link = strings.Split(line, "-")
            if _, ok := core.ColonyInstance.RoomTunnels[Link[0]]; !ok {
                core.HandleDataError("room " + Link[1] + " is linked to unknow room " + Link[0])
            } else if _, ok := core.ColonyInstance.RoomTunnels[Link[1]]; !ok {
                core.HandleDataError("room " + Link[0] + " is linked to unknow room " + Link[1])
            } else {
                if Link[0] == Link[1] {
                    core.HandleDataError("room " + Link[0] + " is linked to itself")
                }
                DuplicatedTunnels(Link[0], Link[1])
                core.ColonyInstance.RoomTunnels[Link[0]] = append(core.ColonyInstance.RoomTunnels[Link[0]], Link[1])
                core.ColonyInstance.RoomTunnels[Link[1]] = append(core.ColonyInstance.RoomTunnels[Link[1]], Link[0])
            }
        } else if line == "" || strings.HasPrefix(line, "#") {
            continue
        } else {
            RoomName = strings.Split(line, " ")
            if _, ok := core.ColonyInstance.RoomTunnels[RoomName[0]]; ok {
                core.HandleDataError("room " + RoomName[0] + " is duplicated")
            } else {
                core.ColonyInstance.RoomTunnels[RoomName[0]] = make([]string, 0)
            }
        }
    }

	if err := scanner.Err(); err != nil {
		core.HandleError(err.Error())
	}
    ValidateData(NumOfStartRoom, NumOfEndroom)
    PrintFileContent(filename)
}

func ValidateData(NumOfStartRoom int, NumOfEndroom int) {
    if NumOfStartRoom != 1 {
        if NumOfStartRoom > 1 {
            core.HandleDataError("more than one start room")
        } else if NumOfStartRoom < 1 {
            core.HandleDataError("no start room found")
        }
    }

    if NumOfEndroom != 1 {
        if NumOfEndroom > 1 {
            core.HandleDataError("more than one end room")
        } else if NumOfEndroom < 1 {
            core.HandleDataError("no end room found")
        }
    }
}

func DuplicatedTunnels(room1 string, room2 string) {
    for _, room := range core.ColonyInstance.RoomTunnels[room1] {
        if room == room2 {
            core.HandleDataError("two tunnels joining room " + room1 + " and room " + room2)
        }
    }
}

func PrintFileContent(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        core.HandleError(err.Error())
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "##") {
            fmt.Println(line)
        } else if line == "" || strings.HasPrefix(line, "#") {
            continue
        } else {
            fmt.Println(line)
        }
    }

	if err := scanner.Err(); err != nil {
		core.HandleError(err.Error())
	}
}
