package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

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

// implementation of deep first search
// adjacencyList - explain entire graph
// currentVertex - current location in graph
// listOfVisited - list where marked visited vertexes, empty in the start
func DFS(adjacencyList map[string][]string, currentVertex string, listOfVisited []string) []string {
	listOfVisited = append(listOfVisited, currentVertex)
	childs := adjacencyList[currentVertex]
	if len(childs) == 0 {
		return listOfVisited
	}

	for _, child := range childs {
		if !isContain(listOfVisited, child) {
			listOfVisited = DFS(adjacencyList, child, listOfVisited)
		}
	}

	return listOfVisited
}

// checks whether the array contains an element
func isContain(arr []string, target string) bool {
	if len(arr) == 0 {
		return false
	}

	for _, elem := range arr {
		if elem == target {
			return true
		}
	}

	return false
}

type Farm struct {
	antAmount string
	start     string
	end       string
	links     []string
	rooms     []string
}

type Vertex struct {
	Name string
	X    int
	Y    int
}

type UpdatedFarm struct {
	AntAmount     int
	Start         Vertex
	End           Vertex
	AdjacencyList map[string][]string
	Rooms         map[string]Vertex
	Weights       map[[2]string]float64
}

// return empty struct in error
func TransformToVertex(data string) Vertex {
	var result Vertex
	splittedData := strings.Split(data, " ")
	if len(splittedData) == 0 {
		return Vertex{}
	}

	result.Name = splittedData[0]

	var err error
	result.X, err = strconv.Atoi(splittedData[1])
	if err != nil {
		return Vertex{}
	}

	result.Y, err = strconv.Atoi(splittedData[2])
	if err != nil {
		return Vertex{}
	}

	return result
}

func UpdateFarm(raw_farm Farm) (result UpdatedFarm, err error) {
	result.AntAmount, err = strconv.Atoi(raw_farm.antAmount)
	if err != nil {
		return result, fmt.Errorf("UpdateFarm: %w", err)
	}

	var emptyV Vertex
	result.Start = TransformToVertex(raw_farm.start)
	if result.Start == emptyV {
		return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform start field of Farm")
	}

	result.End = TransformToVertex(raw_farm.end)
	if result.Start == emptyV {
		return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform end field of Farm")
	}

	result.AdjacencyList = TransformToAdjacencyList(raw_farm.links)
	if len(result.AdjacencyList) == 0 {
		return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform links field of Farm")
	}

	result.Rooms = make(map[string]Vertex)
	for _, elem := range raw_farm.rooms {
		var newVertex Vertex = TransformToVertex(elem)
		if newVertex == emptyV {
			return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform rooms field of Farm")
		}

		result.Rooms[newVertex.Name] = newVertex
	}

	result.Weights = make(map[[2]string]float64)

	return result, nil
}

// Calculates distances between vertexes
// and saves them in map
func CalculateWeights(farm *UpdatedFarm, startPoint Vertex, visited []string) {
	visited = append(visited, startPoint.Name)

	for _, child := range farm.AdjacencyList[startPoint.Name] {
		_, ok := farm.Weights[[2]string{startPoint.Name, child}]
		if ok {
			continue
		}

		var tempVertex Vertex
		if child == farm.Start.Name {
			tempVertex = farm.Start
		} else if child == farm.End.Name {
			tempVertex = farm.End
		} else {
			tempVertex = farm.Rooms[child]
		}

		distance := DistanceBetweenVertex(startPoint, tempVertex)
		farm.Weights[[2]string{startPoint.Name, tempVertex.Name}] = distance
		farm.Weights[[2]string{tempVertex.Name, startPoint.Name}] = distance

		if !isContain(visited, child) {
			CalculateWeights(farm, tempVertex, visited)
		}

	}

}

// Возвращает максимально возможное число
// непересекающихся путей к конечной точке
func CountPaths(farm UpdatedFarm) int {
	return len(farm.AdjacencyList[farm.End.Name])
}

func main() {
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

	farm := Farm{
		antAmount: "9",
		start:     "start 0 3",
		end:       "end 10 1",
		links: []string{"G0-G1", "G1-G2", "G2-G3", "G3-G4", "G4-D3", "start-A0", "A0-A1", "A0-D1", "A1-A2",
			"A1-B1", "A2-end", "A2-C3", "start-B0", "B0-B1", "B1-E2", "start-C0", "C0-C1", "C1-C2",
			"C2-C3", "C3-I4", "D1-D2", "D1-F2", "D2-E2", "D2-D3", "D2-F3", "D3-end", "F2-F3", "F3-F4",
			"F4-end", "I4-I5", "I5-end"},
		rooms: []string{"C0 1 0", "C1 2 0", "C2 3 0", "C3 4 0", "I4 5 0", "I5 6 0", "A0 1 2", "A1 2 1",
			"A2 4 1", "B0 1 4", "B1 2 4", "E2 6 4", "D1 6 3", "D2 7 3", "D3 8 3", "H4 4 2", "H3 5 2",
			"F2 6 2", "F3 7 2", "F4 8 2", "G0 1 5", "G1 2 5", "G2 3 5", "G3 4 5", "G4 6 5"},
	}

	newFarm, err := UpdateFarm(farm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//fmt.Println(newFarm)
	// visited = DFS(newFarm.AdjacencyList, "1", []string{})
	//fmt.Println(visited)

	CalculateWeights(&newFarm, newFarm.Start, []string{})
	// fmt.Println(newFarm.Weights)

	for i := 0; i < CountPaths(newFarm); i++ {
		path := DijkstraAlgo(&newFarm)

		fmt.Println(path)
	}

}

// evaluate distance between two vertexes
func DistanceBetweenVertex(first Vertex, second Vertex) float64 {
	result := math.Sqrt(math.Pow(float64(first.X-second.X), 2) + math.Pow(float64(first.Y-second.Y), 2))
	return result
}

// Algorithm find shortest path in graph
// between start and end
// result path described like array of Vertexes
func DijkstraAlgo(farm *UpdatedFarm) []Vertex {
	var result []Vertex

	var visited []string
	currentVert := farm.Start
	result = append(result, currentVert)

	for {
		visited = append(visited, currentVert.Name)
		minDistance := math.Inf(1)

		var nextVertex Vertex
		for _, child := range farm.AdjacencyList[currentVert.Name] {
			// пропускаем маршруты назад
			if isContain(visited, child) {
				continue
			}

			distance := farm.Weights[[2]string{currentVert.Name, child}]
			if distance < minDistance {
				minDistance = distance
				nextVertex = farm.Rooms[child]
			}

			// если следующий пункт конечный, то
			// сразу добавляем его и выходим из цикла
			if child == farm.End.Name {
				result = append(result, farm.End)
				nextVertex = farm.End
				break
			}
		}

		// меняем весы для поиска следующего пути
		// в следующей итерации данная вершина не будет выбрана
		for _, elem := range farm.AdjacencyList[nextVertex.Name] {
			farm.Weights[[2]string{elem, nextVertex.Name}] = math.Inf(1)
			farm.Weights[[2]string{nextVertex.Name, elem}] = 0
		}

		// проверка на то, что мы достигли конечный пункт
		if nextVertex.Name == farm.End.Name {
			break
		}

		result = append(result, nextVertex)
		currentVert = nextVertex
	}

	return result
}
