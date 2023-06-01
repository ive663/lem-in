package main

import "fmt"

type Farm struct {
	antAmount int
	queue     []int
	paths     [][]string
}

func main() {
	// farm := Farm{
	// 	antAmount: 10,
	// 	paths: [][]string{[]string{"t", "E", "a", "m", "end"},
	// 		[]string{"h", "A", "c", "k", "end"},
	// 		[]string{"0", "o", "n", "e", "end"},
	// 	},
	// 	queue: []int{4, 3, 3},
	// }

	farm := Farm{
		antAmount: 20,
		queue:     []int{11, 9},
		paths:     [][]string{[]string{"3"}, []string{"1", "2", "3"}},
	}

	Printer(&farm)
}

// func that prints result of lem-in project
func Printer(farm *Farm) {
	var antsPosition []int
	FillWitValue(&antsPosition, -1, farm.antAmount)

	pathsNum := len(farm.paths)
	edge := pathsNum
	startPos := 0
	for {

		if startPos == farm.antAmount-1 {
			break
		}

		for antInd := startPos; antInd < edge; antInd++ {
			antsPosition[antInd] += 1
			PrintPosition(farm, &startPos, antInd, antsPosition[antInd])
		}

		fmt.Println()

		if edge+pathsNum <= farm.antAmount {
			edge += pathsNum
		} else {
			edge = farm.antAmount
		}

	}
}

func PrintPosition(farm *Farm, startPos *int, antInd, antPosition int) {
	pathNum := antInd % len(farm.paths)
	pathLen := len(farm.paths[pathNum])

	if antPosition < pathLen {
		fmt.Printf("L%d-%s ", antInd+1, farm.paths[pathNum][antPosition])
	} else {
		*startPos = antInd + 1
	}

}

func FillWitValue(arr *[]int, value, amount int) {
	for i := 0; i < amount; i++ {
		*arr = append(*arr, value)
	}
}
