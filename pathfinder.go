package main

import (
	"math"
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

func DistanceBetweenVertex(first Vertex, second Vertex) float64 {
	xAxes := math.Pow(float64(first.X), 2) - math.Pow(float64(second.X), 2)
	yAxes := math.Pow(float64(first.Y), 2) - math.Pow(float64(second.Y), 2)
	return math.Sqrt(xAxes + yAxes)
}

// Algorithm find shortest path in graph
// between start and end
// result path described like array of Vertexes
func DijkstraAlgo(farm Farm) ([]Vertex, error) {
	var result []Vertex

	currentVert := farm.Start
	result = append(result, currentVert)

	for currentVert != farm.End {
		minDistance := math.Inf(1)
		var nextVertex Vertex
		for _, child := range farm.AdjacencyList[currentVert.Name] {
			tempVertex := farm.Rooms[child]

			distance := DistanceBetweenVertex(currentVert, tempVertex)
			// if distance < 0 {
			// 	return nil, fmt.Errorf("DijkstraAlgo: can't evaluate distance between vertexes(%f): %s and %s",
			// 		distance, currentVert.Name, tempVertex.Name)
			// }

			if child == farm.End.Name {
				result = append(result, farm.End)
				nextVertex = Vertex{}
				break
			} else if distance < minDistance {
				minDistance = distance
				nextVertex = tempVertex
			}

		}
		if nextVertex.Name != "" {
			result = append(result, nextVertex)
		} else {
			break
		}

		currentVert = nextVertex
	}

	return result, nil
}

// func main() {
// 	farm := Farm{
// 		antAmount: "3",
// 		start:     "1 23 3",
// 		end:       "0 9 5",
// 		links: []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5",
// 			"4-2", "2-1", "7-6", "7-2", "7-4", "6-5"},
// 		rooms: []string{"2 16 7", "3 16 3", "4 16 5", "5 9 3", "6 1 5", "7 4 8"},
// 	}

// 	newFarm, err := UpdateFarm(farm)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	//fmt.Println(newFarm)
// 	// visited := DFS(newFarm.AdjacencyList, "1", []string{})
// 	// fmt.Println(visited)
// 	path, err := DijkstraAlgo(newFarm)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	fmt.Println(path)
// }
