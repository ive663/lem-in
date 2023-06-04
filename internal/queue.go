package internal

import "fmt"

func Queue(paths [][]string, antsAmount int) []int {

	antsQueue := make([]int, len(paths))
	rooms := make([]int, len(paths))

	for i := range paths {
		rooms[i] = len(paths[i])
	}

	p := 1
	for antsAmount > 0 {
		fmt.Println("rooms: ", rooms)
		fmt.Println("paths: ", len(paths))
		fmt.Println("p: ", p)
		fmt.Println("antsQueue: ", antsQueue)
		fmt.Println("antsAmount: ", antsAmount)

		if p == len(paths) {
			antsQueue[p-1]++
			antsAmount--
			p = 1
		}
		fmt.Println("p: ", p)

		if rooms[p-1]+antsQueue[p-1] <= rooms[p]+antsQueue[p] {
			antsQueue[p-1]++
			antsAmount--
		} else if rooms[p-1]+antsQueue[p-1] > rooms[p]+antsQueue[p] {
			p++
		}
	}
	// fmt.Println(antsQueue)
	return antsQueue
}

// func Queue(paths [][]string, antsAmount int) [10]int {

// paths1 := [][]string{
// 	{"h", "A"},
// 	{"h", "A", "t"},
// 	{"h", "A", "t"},
// 	{"0", "o", "h", "e", "h", "A", "t", "A"},
// }

// paths2 := [][]string{
// 	{"h"},
// 	{"h", "A"},
// 	{"h", "A", "t"},
// }

// 	// paths3 := [][]string{
// 	// 	{"t", "E"},
// 	// 	{"h", "A", "c"},
// 	// 	{"0", "o", "h"},
// 	// 	{"e", "E", "a", "m", "A", "c", "k"},
// 	// }

// 	// antsAmount = 10
// 	rooms := []int{}
// 	for i, _ := range paths {
// 		rooms = append(rooms, len(paths[i]))
// 	}

// 	// fmt.Println("rooms: ", rooms)
// 	// fmt.Println("len path3: ", len(paths2[2]))
// 	// fmt.Println("len rooms: ", len(rooms))

// 	antsQueue := [10]int{}
// 	p := 1
// 	for antsAmount > 0 {
// 		if len(rooms) == p {
// 			// fmt.Println("p: ", p)
// 			antsQueue[p-1]++
// 			// fmt.Println("antsQueue: ", antsQueue)
// 			antsAmount--
// 			p = 1
// 		}

// 		if rooms[p-1]+antsQueue[p-1] <= rooms[p]+antsQueue[p] {
// 			antsQueue[p-1]++
// 			antsAmount--
// 			// fmt.Println("rooms:     ", rooms)

// 			// fmt.Println("antsQueue: ", antsQueue)
// 		} else if rooms[p-1]+antsQueue[p-1] > rooms[p]+antsQueue[p] {
// 			// fmt.Println("p>", p)
// 			p++
// 		}
// 	}
// 	// fmt.Println("antsAmount: ", antsAmount)
// 	// fmt.Println("res queue: ", rooms)
// 	// fmt.Println("res antsQueue: ", antsQueue)
// 	return antsQueue
// }
