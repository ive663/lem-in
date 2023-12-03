package internal

import "fmt"

// распредиление муровьев по направлению в очередь...
func Queue(paths [][]string, antsAmount int) []int {

	antsQueue := make([]int, len(paths))
	rooms := make([]int, len(paths))

	for i := range paths {
		rooms[i] = len(paths[i])
	}

	for antsAmount > 0 {
		fmt.Println("rooms: ", rooms)
		fmt.Println("paths: ", len(paths))
		fmt.Println("antsQueue: ", antsQueue)
		fmt.Println("antsAmount: ", antsAmount)

		indexOfInsert := checkLowestPath(rooms, antsQueue)
		antsQueue[indexOfInsert] += 1
		antsAmount -= 1
	}
	return antsQueue
}

// проверка наименьшего количества муровьев в очерреди...
func checkLowestPath(rooms []int, antsQueue []int) int {
	lowestValue := 10000
	lowestInd := 0
	for indOfPath := range rooms {
		// summ of rooms and ants in one path
		sum := rooms[indOfPath] + antsQueue[indOfPath]
		if sum < lowestValue {
			lowestValue = sum
			lowestInd = indOfPath
		}
	}

	return lowestInd
}
