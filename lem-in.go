package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// frm := &Farm{}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter input file")
		os.Exit(0)
	}
	filename := args[0] //obtaining farm's filename from os.Args

	farm, err := prepFarm(filename)
	check(err)

	fmt.Println("=================================")
	// fmt.Println("ALL: ", frm)
	fmt.Println("Start: ", farm.Start)
	fmt.Println("End: ", farm.End)
	fmt.Println("antAmunt: ", farm.AntAmount)
	fmt.Println("links: ", farm.AdjacencyList)
	fmt.Println("rooms: ", farm.Rooms)
	moveAnts()
	// paths := path(links, start, end)
	// fmt.Println(paths)
}

type Farm struct {
	AntAmount     int
	Start         Vertex
	End           Vertex
	AdjacencyList map[string][]string
	Rooms         map[string]Vertex
	bestPaths     Path
}

type Vertex struct {
	Name string
	X    int
	Y    int
}
type Path struct {
	Vertices [][]Vertex
}
type Room struct {
  pathNbr int
  isBusy int
  location Vertex
}

type Queue struct{
  rooms [][]Room
}

func fillQueueFromPath(path Path) Queue {
	queue := Queue{}

	for _, vertices := range path.Vertices {
		roomList := make([]Room, 0)

		for _, vertex := range vertices {
			room := Room{
				pathNbr: 0, // Set the initial path number, modify as needed
				isBusy: 0, // Set the initial busy status, modify as needed
				location: vertex,
			}
			roomList = append(roomList, room)
		}

		queue.rooms = append(queue.rooms, roomList)
	}
  fmt.Println(">>>>>>", queue)
	return queue
}




var pathN int
var nbrOfRoomsInPath int
var nbrOfAntsInPath int




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

func moveAnts() {
  // antsN1 := 10 
	path1 := Path{
		Vertices: [][]Vertex{
			{{"t", 1, 9}, {"E", 5, 9}, {"a", 8, 9}, {"m", 8, 6}},
			{{"h", 4, 6}, {"A", 5, 2}, {"c", 8, 1}, {"k", 11, 2}},
			{{"0", 4, 8}, {"o", 6, 8}, {"h", 4, 6}, {"e", 8, 4}},
		},
	}

  // antsN2 := 3 
 //  path2 := Path{
	// 	Vertices: [][]Vertex{
	// 		{{"t", 1, 9}},
	// 		{{"h", 4, 6}, {"A", 5, 2}},
	// 		{{"0", 4, 8}, {"o", 6, 8}, {"h", 4, 6}, {"e", 8, 4}},
	// 	},
	// }
  // antsN3 := 6 
  path3 := Path{
		Vertices: [][]Vertex{
			{{"t", 1, 9}, {"E", 5, 9}},
			{{"h", 4, 6}, {"A", 5, 2}, {"c", 8, 1}},
			{{"0", 4, 8}, {"o", 6, 8}, {"h", 4, 6}},
      {{"e", 8, 4}, {"E", 5, 9}, {"a", 8, 9}, {"m", 8, 6}, {"A", 5, 2}, {"c", 8, 1}, {"k", 11, 2}},
		},
	}

  fillQueueFromPath(path3)



	firstPath := path1.Vertices[0]
	firstVertex := path1.Vertices[0][0]
	fmt.Println(firstPath)
	fmt.Println(firstVertex)
  // fmt.Println("<>",r)
}

func prepFarm(filename string) (result Farm, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return Farm{}, fmt.Errorf("prepFarm: %w", err)
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
		return Farm{}, fmt.Errorf("prepFarm: empty file")
	}

	result.AntAmount, err = strconv.Atoi(fileLines[0])
	if err != nil {
		return Farm{}, fmt.Errorf("prepFarm: failed to parse ant amount: %w", err)
	}

	// var emptyV Vertex
	var links []string
	var rooms []string
	var start string
	var end string

	for i, line := range fileLines[1:] {
		switch i {
		case idxOfStart, idxOfEnd:
			if i == idxOfStart {
				start = line
			} else {
				end = line
			}
		default:
			if strings.Contains(line, "-") {
				links = append(links, line)
			} else if !strings.Contains(line, "#") {
				rooms = append(rooms, line)
			}
		}
	}

	result.Start, err = TransformToVertex(start)
	if err != nil {
		return Farm{}, fmt.Errorf("prepFarm: failed to transform start field: %w", err)
	}

	result.End, err = TransformToVertex(end)
	if err != nil {
		return Farm{}, fmt.Errorf("prepFarm: failed to transform end field: %w", err)
	}

	result.AdjacencyList = TransformToAdjacencyList(links)
	if len(result.AdjacencyList) == 0 {
		return Farm{}, fmt.Errorf("prepFarm: failed to transform links field")
	}

	result.Rooms = make(map[string]Vertex)
	for _, elem := range rooms {
		fmt.Println("elem: ", elem)
		if elem == "" {
			continue
		}
		newVertex, err := TransformToVertex(elem)
		if err != nil {
			return Farm{}, fmt.Errorf("prepFarm: failed to transform rooms field: %w", err)
		}
		result.Rooms[newVertex.Name] = newVertex
	}

	return result, nil
}

func TransformToVertex(data string) (Vertex, error) {
	parts := strings.Fields(data)
	if len(parts) < 3 {
		return Vertex{}, fmt.Errorf("TransformToVertex: invalid data")
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return Vertex{}, fmt.Errorf("TransformToVertex: %w", err)
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return Vertex{}, fmt.Errorf("TransformToVertex: %w", err)
	}
	return Vertex{
		Name: parts[0],
		X:    x,
		Y:    y,
	}, nil
}

// 	// // var emptyV Vertex
// 	// var links []string
// 	// var rooms []string
// 	// var start string
// 	// var end string

// 	// for i, line := range fileLines[1:] {
// 	// 	switch i {
// 	// 	case idxOfStart, idxOfEnd:
// 	// 		if i == idxOfStart {
// 	// 			start = line
// 	// 		} else {
// 	// 			end = line
// 	// 		}
// 	// 	default:
// 	// 		if strings.Contains(line, "-") {
// 	// 			links = append(links, line)
// 	// 		} else if !strings.Contains(line, "#") {
// 	// 			rooms = append(rooms, line)
// 	// 		}
// 	// 	}
// 	// }

// 	antAmount, err := strconv.Atoi(fileLines[0])
// 	if err != nil {
// 		return result, fmt.Errorf("prepFarm:Atoi: %w", err)
// 	}

// 	farm := Farm{
// 		AntAmount: antAmount,
// 	}

// 	farm.Start, err = TransformToVertex(start)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}

// 	farm.End, err = TransformToVertex(end)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}

// 	// if farm.Start == emptyV {
// 	// 	return Farm{}, fmt.Errorf("UpdateFarm: can't transform start field of Farm")
// 	// }
// 	// if farm.End == emptyV {
// 	// 	return Farm{}, fmt.Errorf("UpdateFarm: can't transform end field of Farm")
// 	// }
// 	farm.AdjacencyList = make(map[string][]string)
// 	farm.Rooms = make(map[string]Vertex)

// 	for _, line := range fileLines[1:] {
// 		if strings.Contains(line, "-") {
// 			parts := strings.Split(line, "-")
// 			if len(parts) != 2 {
// 				continue
// 			}
// 			from, err1 := TransformToVertex(parts[0])
// 			to, err2 := TransformToVertex(parts[1])
// 			if err1 != nil || err2 != nil {
// 				continue
// 			}
// 			farm.AdjacencyList[from.Name] = append(farm.AdjacencyList[from.Name], to.Name)
// 			farm.AdjacencyList[to.Name] = append(farm.AdjacencyList[to.Name], from.Name)
// 		} else if !strings.HasPrefix(line, "#") {
// 			vertex, err := TransformToVertex(line)
// 			if err != nil {
// 				continue
// 			}
// 			farm.Rooms[vertex.Name] = vertex
// 		}
// 	}
// 	if len(farm.AdjacencyList) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: no links found")
// 	}

// 	if len(farm.Rooms) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: no rooms found")
// 	}

// 	fmt.Println(fileLines)
// 	return farm, nil
// }

// // func path(links []string, start string, end string) [][]string {
// // 	p := make([][]string, 1)
// // 	startloc := strings.Split(start, " ")
// // 	// counter := 0
// // 	p[0] = append(p[0], startloc[0])
// // 	tmp := []string{}
// // 	// for i1 := 0; i1 <= counter; i1++ {
// // 	for i, val := range links {
// // 		if strings.Contains(val, p[i][i]) {

// // 			tmp = append(tmp, links[i])

// func remove(l []string, item string) []string {
// 	for i, other := range l {
// 		if other == item {
// 			return append(l[:i], l[i+1:]...)
// 		}
// 	}
// 	return l
// }

// func set(list) {
// 	antAmount := list[0]
// }

// func (in *Input) prep(in *Input) error {
// 	f, err := os.Open(filename)

// 	check(err)

// 	fileScanner := bufio.NewScanner(f)
// 	fileScanner.Split(bufio.ScanLines)

// 	var counter int
// 	var fileLines []string
// 	var idxOfStart int
// 	var idxOfEnd int

// 	for fileScanner.Scan() {
// 		if fileScanner.Text() == "##start" {
// 			idxOfStart = counter + 1
// 		}
// 		if fileScanner.Text() == "##end" {
// 			idxOfEnd = counter + 1
// 		}
// 		counter++
// 		fileLines = append(fileLines, fileScanner.Text())
// 	}
// 	f.Close()

// 	inp := &Input{}

// 	var links []string
// 	var rooms []string
// 	for i, line := range fileLines {
// 		if i == 0 {
// 			i.antAmount = line
// 		} else if idxOfStart == i {
// 			i.start = line
// 		} else if idxOfEnd == i {
// 			i.end = line
// 		} else if strings.Contains(line, "-") {
// 			in.links = append(links, line)
// 		} else if line != "##start" && line != "##end" {
// 			in.rooms = append(rooms, line)
// 		}
// 		fmt.Println(line)
// 	}

// 	fmt.Println(fileLines)
// 	return in
// }

// func prepFarm(filename string) (Farm, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: failed to open file: %w", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)

// 	var lines []string
// 	var start, end string

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "##start" {
// 			start = lines[len(lines)-1]
// 		}
// 		if line == "##end" {
// 			end = lines[len(lines)-1]
// 		}
// 		lines = append(lines, line)
// 	}

// 	if len(lines) < 2 {
// 		return Farm{}, fmt.Errorf("prepFarm: file is empty")
// 	}

// 	antAmount, err := strconv.Atoi(lines[0])
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}

// 	farm := Farm{
// 		AntAmount: antAmount,
// 	}

// 	farm.Start, err = TransformToVertex(start)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}

// 	farm.End, err = TransformToVertex(end)
// 	if err != nil {
// 		return Farm{}, fmt.Errorf("prepFarm: %w", err)
// 	}

// 	farm.AdjacencyList = make(map[string][]string)
// 	farm.Rooms = make(map[string]Vertex)

// 	for _, line := range lines[1:] {
// 		if strings.Contains(line, "-") {
// 			parts := strings.Split(line, "-")
// 			if len(parts) != 2 {
// 				continue
// 			}
// 			from, err1 := TransformToVertex(parts[0])
// 			to, err2 := TransformToVertex(parts[1])
// 			if err1 != nil || err2 != nil {
// 				continue
// 			}
// 			farm.AdjacencyList
// 			}
// 			farm.Rooms[vertex.Name] = vertex
// 		}
// 	}

// 	if len(farm.AdjacencyList) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: no links found")
// 	}

// 	if len(farm.Rooms) == 0 {
// 		return Farm{}, fmt.Errorf("prepFarm: no rooms found")
// 	}

// 	return farm, nil
// }

func check(e error) {
	if e != nil {
		log.Println("err:", e)
		return

	}
}
