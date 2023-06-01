package internal

// "log"

// // function that transform list of edges to adjacency list
// func TransformToAdjacencyList(listOfEdges []string) map[string][]string {
// 	var result map[string][]string = make(map[string][]string)

// 	if len(listOfEdges) == 0 {
// 		return result
// 	}

// 	for _, pairOfVertex := range listOfEdges {
// 		vertexes := strings.Split(pairOfVertex, "-")
// 		result[vertexes[0]] = append(result[vertexes[0]], vertexes[1])
// 		result[vertexes[1]] = append(result[vertexes[1]], vertexes[0])
// 	}

// 	return result
// }

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

// func UpdateFarm(raw_farm Farm) (result UpdatedFarm, err error) {
// 	result.AntAmount, err = strconv.Atoi(raw_farm.antAmount)
// 	if err != nil {
// 		return result, fmt.Errorf("UpdateFarm: %w", err)
// 	}

// 	result.Start = GetName(raw_farm.start)

// 	result.End = GetName(raw_farm.end)

// 	result.AdjacencyList = TransformToAdjacencyList(raw_farm.links)
// 	if len(result.AdjacencyList) == 0 {
// 		return UpdatedFarm{}, fmt.Errorf("UpdateFarm: can't transform links field of Farm")
// 	}

// 	result.Weights = make(map[[2]string]bool)

// 	return result, nil
// }

// func GetName(info string) string {
// 	splittedData := strings.Split(info, " ")
// 	if len(splittedData) == 0 {
// 		return ""
// 	}

// 	return splittedData[0]
// }

// Создает начальную схему весов, где все равно 0
func CalculateWeights(farm *UpdatedFarm, startPoint string, visited []string) {
	visited = append(visited, startPoint)

	for _, child := range farm.AdjacencyList[startPoint] {
		_, ok := farm.Weights[[2]string{startPoint, child}]
		if ok {
			continue
		}

		farm.Weights[[2]string{startPoint, child}] = true
		farm.Weights[[2]string{child, startPoint}] = true

		if !isContain(visited, child) {
			CalculateWeights(farm, child, visited)
		}

	}

}

// Возвращает максимально возможное число
// непересекающихся путей к конечной точке
func CountPaths(farm UpdatedFarm) int {
	endPointsNum := len(farm.AdjacencyList[farm.End])
	startPointsNum := len(farm.AdjacencyList[farm.Start])

	if endPointsNum < startPointsNum {
		return endPointsNum
	} else {
		return startPointsNum
	}
}

func ClosePaths(farm *UpdatedFarm, child string) {
	parents := farm.AdjacencyList[child]

	for _, parent := range parents {
		farm.Weights[[2]string{parent, child}] = false
	}
}

// Algorithm find shortest path in graph
// between start and end
// result path described like array of Vertexes
func DijkstraAlgo(farm *UpdatedFarm, parent string, parents []string) []string {
	var result []string

	childs := farm.AdjacencyList[parent]

	// поиск конца среди текущих потомков
	for _, child := range childs {
		if child == farm.End {
			result = append(result, parents...)
			result = append(result, parent)
			result = append(result, child)
			return result
		}
	}

	// переход к потомкам на уровень ниже
	for _, child := range childs {
		if !isContain(parents, child) &&
			farm.Weights[[2]string{parent, child}] {
			parents = append(parents, parent)
			result = DijkstraAlgo(farm, child, parents)
		}

		// если нет верного пути впереди
		if result == nil {
			continue
		}

		// делаем так чтобы текущая вершина не была выбрана
		// из родителей
		ClosePaths(farm, child)

		return result
	}

	return nil
}
