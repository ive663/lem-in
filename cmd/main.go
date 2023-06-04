package main

import (
	"fmt"
	"os"

	"github.com/ive663/lem-in/internal"
)

var Old_farm = internal.Farm{
	AntAmount: "9",
	Start:     "start 0 3",
	End:       "end 10 1",
	Links: []string{"G0-G1", "G1-G2", "G2-G3", "G3-G4", "G4-D3", "start-A0", "A0-A1", "A0-D1", "A1-A2",
		"A1-B1", "A2-end", "A2-C3", "start-B0", "B0-B1", "B1-E2", "start-C0", "C0-C1", "C1-C2",
		"C2-C3", "C3-I4", "D1-D2", "D1-F2", "D2-E2", "D2-D3", "D2-F3", "D3-end", "F2-F3", "F3-F4",
		"F4-end", "I4-I5", "I5-end"},
	Rooms: []string{"C0 1 0", "C1 2 0", "C2 3 0", "C3 4 0", "I4 5 0", "I5 6 0", "A0 1 2", "A1 2 1",
		"A2 4 1", "B0 1 4", "B1 2 4", "E2 6 4", "D1 6 3", "D2 7 3", "D3 8 3", "H4 4 2", "H3 5 2",
		"F2 6 2", "F3 7 2", "F4 8 2", "G0 1 5", "G1 2 5", "G2 3 5", "G3 4 5", "G4 6 5"},
}

func main() {

	input := os.Args[1]
	new_farm, err := internal.PrepFarm(input)
	if err != nil {
		fmt.Errorf("new farm err %w", err)
		return
	}

	// _, err := internal.UpdateFarm(Old_farm)
	// if err != nil {
	// 	fmt.Errorf("Updated farm err %w", err)
	// 	return
	// }
	fmt.Println("=================================")
	// fmt.Println("ALL: ", frm)
	fmt.Println("Start: ", new_farm.Start)
	fmt.Println("End: ", new_farm.End)
	fmt.Println("antAmunt: ", new_farm.AntAmount)
	fmt.Println("AdjacencyList: ", new_farm.AdjacencyList)
	fmt.Println()
	internal.Queue()

	// paths := path(links, start, end)
	// fmt.Println(paths)
}

// farm := Farm{
// 	antAmount: "10",
// 	start:     "start 1 6",
// 	end:       "end 11 6",
// 	links: []string{"start-t", "n-e", "a-m", "A-c", "0-o", "E-a", "k-end", "start-h", "o-n",
// 		"m-end", "t-E", "start-0", "h-A", "e-end", "c-k", "n-m", "h-n"},
// 	rooms: []string{"0 4 8", "o 6 8", "n 6 6", "e 8 4", "t 1 9", "E 5 9", "a 8 9", "m 8 6",
// 		"h 4 6", "A 5 2", "c 8 1", "k 11 2"},
// }

// farm := Farm{
// 	antAmount: "4",
// 	start:     "0 0 3",
// 	end:       "1 8 3",
// 	links:     []string{"0-2", "2-3", "3-1"},
// 	rooms:     []string{"2 2 5", "3 4 0"},
// }

// farm := Farm{
// 	antAmount: "9",
// 	start:     "start 0 3",
// 	end:       "end 10 1",
// 	links: []string{"G0-G1", "G1-G2", "G2-G3", "G3-G4", "G4-D3", "start-A0", "A0-A1", "A0-D1", "A1-A2",
// 		"A1-B1", "A2-end", "A2-C3", "start-B0", "B0-B1", "B1-E2", "start-C0", "C0-C1", "C1-C2",
// 		"C2-C3", "C3-I4", "D1-D2", "D1-F2", "D2-E2", "D2-D3", "D2-F3", "D3-end", "F2-F3", "F3-F4",
// 		"F4-end", "I4-I5", "I5-end"},
// 	rooms: []string{"C0 1 0", "C1 2 0", "C2 3 0", "C3 4 0", "I4 5 0", "I5 6 0", "A0 1 2", "A1 2 1",
// 		"A2 4 1", "B0 1 4", "B1 2 4", "E2 6 4", "D1 6 3", "D2 7 3", "D3 8 3", "H4 4 2", "H3 5 2",
// 		"F2 6 2", "F3 7 2", "F4 8 2", "G0 1 5", "G1 2 5", "G2 3 5", "G3 4 5", "G4 6 5"},
// }

// newFarm, err := UpdateFarm(farm)
// if err != nil {
// 	log.Fatalln(err.Error())
// }

// //fmt.Println(newFarm)
// // visited = DFS(newFarm.AdjacencyList, "1", []string{})
// //fmt.Println(visited)

// CalculateWeights(&newFarm, newFarm.Start, []string{})
// //fmt.Println(newFarm.Weights)

// numOfPaths := CountPaths(newFarm)
// paths := [][]string{}
// for len(paths) != numOfPaths {
// 	path := DijkstraAlgo(&newFarm, newFarm.Start, []string{})
// 	fmt.Println(path)

// 	paths = append(paths, path)
// }

///===================================
// type Farm struct {
// 	AntAmount     int
// 	Start         string
// 	End           string
// 	AdjacencyList map[string][]string
// 	Weights       map[[2]string]bool
// }

// paths for test01.txt:
//
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

// // Accessing individual vertices
// firstVertex := path.Vertices[0][0]
// firstPath := path.Vertices[0]
// secondVertex := path.Vertices[1][1]

// paths := [][]string{}
// antsN1 := 10
//  paths1 := [][]string{
//    {"t", "E", "a", "m"},
//   {"h", "A", "c", "k"},
//   {"0", "o", "h", "e"},
// }

// antsN2 := 3
//  paths2 := [][]string{
// 	{"t"},
// 	{"h", "A"},
// 	{"0", "o", "h", "e"},
// }
// antsN3 := 6
//   paths3 := [][]string{
// 			{"t", "E"},
// 			{"h", "A", "c"},
// 			{"0", "o", "h"},
//       {"e", "E", "a", "m", "A", "c", "k"},
// 	}

// 	// firstPath := path1.Vertices[0]
// 	firstVertex := paths3[0][0]
//   maxPathLength := 0

//   for _, p := range paths3{
//     if maxPathLength< len(p) {
//       maxPathLength = len(p)
//     }
//   }
//

//   fmt.Println("maxPathLength: ", maxPathLength)
// 	fmt.Println(firstVertex)
//   // fmt.Println("<>",r)
// }

// func prepFarm(filename string) (result Farm, err error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}
// 	defer f.Close()

// 	fileScanner := bufio.NewScanner(f)
// 	fileScanner.Split(bufio.ScanLines)

// 	var fileLines []string
// 	var idxOfStart, idxOfEnd int
// 	for fileScanner.Scan() {
// 		line := fileScanner.Text()
// 		if line == "##start" {
// 			idxOfStart = len(fileLines)
// 		}
// 		if line == "##end" {
// 			idxOfEnd = len(fileLines)
// 		}
// 		fileLines = append(fileLines, line)
// 	}
// 	if len(fileLines) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: empty file")
// 	}

// 	result.AntAmount, err = strconv.Atoi(fileLines[0])
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: failed to parse ant amount: %w", err)
// 	}

// 	// var emptyV Vertex
// 	var links []string
// 	var rooms []string
// 	var start string
// 	var end string

// 	for i, line := range fileLines[1:] {
// 		switch i {
// 		case idxOfStart, idxOfEnd:
// 			if i == idxOfStart {
// 				start = line
// 			} else {
// 				end = line
// 			}
// 		default:
// 			if strings.Contains(line, "-") {
// 				links = append(links, line)
// 			} else if !strings.Contains(line, "#") {
// 				rooms = append(rooms, line)
// 			}
// 		}
// 	}

// 	result.Start, err = TransformToVertex(start)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: failed to transform start field: %w", err)
// 	}

// 	result.End, err = TransformToVertex(end)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: failed to transform end field: %w", err)
// 	}

// 	result.AdjacencyList = TransformToAdjacencyList(links)
// 	if len(result.AdjacencyList) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: failed to transform links field")
// 	}

// 	result.Rooms = make(map[string]Vertex)
// 	for _, elem := range rooms {
// 		fmt.Println("elem: ", elem)
// 		if elem == "" {
// 			continue
// 		}
// 		newVertex, err := TransformToVertex(elem)
// 		if err != nil {
// 			return Farm{}, fmt.Errorf("prepFarm: failed to transform rooms field: %w", err)
// 		}
// 		result.Rooms[newVertex.Name] = newVertex
// 	}

// 	return result, nil
// }

// func TransformToVertex(data string) (Vertex, error) {
// 	parts := strings.Fields(data)
// 	if len(parts) < 3 {
// 		return Vertex{}, fmt.Errorf("TransformToVertex: invalid data")
// 	}

// 	x, err := strconv.Atoi(parts[1])
// 	if err != nil {
// 		return Vertex{}, fmt.Errorf("TransformToVertex: %w", err)
// 	}
// 	y, err := strconv.Atoi(parts[2])
// 	if err != nil {
// 		return Vertex{}, fmt.Errorf("TransformToVertex: %w", err)
// 	}
// 	return Vertex{
// 		Name: parts[0],
// 		X:    x,
// 		Y:    y,
// 	}, nil
// }
