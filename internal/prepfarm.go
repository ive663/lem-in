package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sorted farm
type UpdatedFarm struct {
	AntAmount     int
	Start         string
	End           string
	AdjacencyList map[string][]string
	Weights       map[[2]string]bool
	Queue         []int
	Paths         [][]string
}

// base farm
type Farm struct {
	AntAmount string
	Start     string
	End       string
	Links     []string
	Rooms     []string
}

// farm parser
func PrepFarm(filename string) (result UpdatedFarm, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return UpdatedFarm{}, fmt.Errorf("prepFarm: %w", err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	var idxOfStart, idxOfEnd int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "##start" {
			idxOfStart = len(fileLines)
		}
		if line == "##end" {
			idxOfEnd = len(fileLines)
		}
		fileLines = append(fileLines, line)
	}
	if len(fileLines) == 0 {
		return UpdatedFarm{}, fmt.Errorf("prepFarm: empty file")
	}

	result.AntAmount, err = strconv.Atoi(fileLines[0])
	if err != nil {
		return UpdatedFarm{}, fmt.Errorf("prepFarm: failed to parse ant amount: %w", err)
	}

	// var emptyV Vertex
	var links []string
	var rooms []string

	for i, line := range fileLines[1:] {
		switch i {
		case idxOfStart, idxOfEnd:
			if i == idxOfStart {
				result.Start = line
			} else {
				result.End = line
			}
		default:
			if strings.Contains(line, "-") {
				links = append(links, line)
			} else if !strings.Contains(line, "#") {
				rooms = append(rooms, line)
			}
		}
	}

	result.AdjacencyList = TransformToAdjacencyList(links)
	if len(result.AdjacencyList) == 0 {
		return UpdatedFarm{}, fmt.Errorf("prepFarm: failed to transform links field")
	}

	return result, nil
}

func UpdateFarm(raw_farm Farm) (result UpdatedFarm, err error) {
	result.AntAmount, err = strconv.Atoi(raw_farm.AntAmount)
	if err != nil {
		return result, fmt.Errorf("UpdateFarm: %w", err)
	}

	result.Start = GetName(raw_farm.Start)

	result.End = GetName(raw_farm.End)

	result.AdjacencyList = TransformToAdjacencyList(raw_farm.Links)
	if len(result.AdjacencyList) == 0 {
		return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform links field of Farm")
	}

	result.Weights = make(map[[2]string]bool)

	return result, nil
}

func GetName(info string) string {
	splittedData := strings.Split(info, " ")
	if len(splittedData) == 0 {
		return ""
	}

	return splittedData[0]
}

// function that transform list of edges to adjacency list
func TransformToAdjacencyList(listOfEdges []string) map[string][]string {
	var result map[string][]string = make(map[string][]string)

	if len(listOfEdges) == 0 {
		return result
	}

	for _, pairOfVertex := range listOfEdges {
		vertexes := strings.Split(pairOfVertex, "-")
		result[vertexes[0]] = append(result[vertexes[0]], vertexes[1])
		result[vertexes[1]] = append(result[vertexes[1]], vertexes[0])
	}

	return result
}

// paths for test01.txt:
// path := Path{
// 	Vertices: [][]Vertex{
// 		{{t, 1, 9}, {E, 5, 9}, {a, 8, 9}, {m, 8, 6}},
// 		{{h, 4, 6}, {A, 5, 2}, {c, 8, 1},{k, 11, 2}},
//  	{{0, 4, 8}, {o, 6, 8}, {h, 4, 6}, {e, 8, 4}},
// 	}

// L1-t L2-h L3-0
// L1-E L2-A L3-o L4-t L5-h L6-0
// L1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0
// L1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t
// L1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E
// L4-end L5-end L6-end L7-m L8-k L9-e L10-a
// L7-end L8-end L9-end L10-m
// L10-end
