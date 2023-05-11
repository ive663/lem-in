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
	farm := Farm{
		antAmount: "3",
		start:     "0 1 0",
		end:       "1 5 0",
		links:     []string{"0-2", "2-3", "3-1"},
		rooms:     []string{"2 9 0", "3 13 0"},
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
		path, visited := DijkstraAlgo(&newFarm, []string{})

		fmt.Println(path)
		fmt.Println(visited)
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
func DijkstraAlgo(farm *UpdatedFarm, visited []string) ([]Vertex, []string) {
	var result []Vertex

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
		farm.Weights[[2]string{currentVert.Name, nextVertex.Name}] = math.Inf(1)
		farm.Weights[[2]string{nextVertex.Name, currentVert.Name}] = 0

		// проверка на то, что мы достигли конечный пункт
		if nextVertex.Name == farm.End.Name {
			break
		}

		result = append(result, nextVertex)
		currentVert = nextVertex
	}

	return result, visited
}
