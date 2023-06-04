package internal

import "fmt"

func Queue(path [][]string, antsAmount int) [][]string {

	paths1 := [][]string{

		{"h", "A"},
		{"h", "A", "t"},
		{"h", "A", "t"},
		{"0", "o", "h", "e", "h", "A", "t", "A"},
	}

	paths2 := [][]string{
		{"h"},
		{"h", "A"},
		{"h", "A", "t"},
	}

	paths3 := [][]string{
		{"t", "E"},
		{"h", "A", "c"},
		{"0", "o", "h"},
		{"e", "E", "a", "m", "A", "c", "k"},
	}

	antsAmount = 10
	rooms := []int{}
	for i, _ := range paths2 {
		rooms = append(rooms, len(paths2[i]))
	}

	fmt.Println("rooms: ", rooms)
	fmt.Println("len path3: ", len(paths2[2]))
	fmt.Println("len rooms: ", len(rooms))

	// var p1 []int
	// var p2 []int
	antsQueue := [3]int{}
	p := 1
	// i := 0
	for antsAmount > 0 {
		if len(rooms) == p {
			fmt.Println("p: ", p)
			antsQueue[p-1]++
			fmt.Println("antsQueue: ", antsQueue)
			antsAmount--

			// fmt.Println("zero:", p)
			p = 1
		}

		if rooms[p-1]+antsQueue[p-1] <= rooms[p]+antsQueue[p] {
			antsQueue[p-1]++
			antsAmount--
			fmt.Println("rooms:     ", rooms)

			fmt.Println("antsQueue: ", antsQueue)
		} else if rooms[p-1]+antsQueue[p-1] > rooms[p]+antsQueue[p] {
			fmt.Println("p>", p)
			p++
		}
		// fmt.Println("p:", p)
	}
	fmt.Println("antsAmount: ", antsAmount)
	fmt.Println("res queue: ", rooms)

	fmt.Println("res antsQueue: ", antsQueue)

}
