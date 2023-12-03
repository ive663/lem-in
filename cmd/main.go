package main

import (
	"fmt"
	"os"

	"github.com/ive663/lem-in/internal"
)

func main() {

	input := os.Args[1]
	new_farm, err := internal.PrepFarm(input)
	if err != nil {
		fmt.Errorf("new farm err %w", err)
		return
	}

	fmt.Println("=================================")
	fmt.Println("Start: ", new_farm.Start)
	fmt.Println("End: ", new_farm.End)
	fmt.Println("antAmunt: ", new_farm.AntAmount)
	fmt.Println("AdjacencyList: ", new_farm.AdjacencyList)
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("path:        ", new_farm.Paths)
	fmt.Println("AntAmaunt:        ", new_farm.AntAmount)
	fmt.Println("queue result: ", internal.Queue(new_farm.Paths, new_farm.AntAmount))

}

// test progoni
// antsAmount1 := 10
// paths1 := [][]string{
// 	{"h"},
// 	{"h", "A"},
// 	{"h", "A", "t"},
// }
// antsAmount2 := 10
// paths2 := [][]string{
// 	{"h", "A"},
// 	{"h", "A", "t"},
// 	{"h", "A", "t"},
// 	{"0", "o", "h", "e", "h", "A", "t", "A"},
// }
// antsAmount3 := 10
// paths3 := [][]string{
// 	{"h", "A"},
// 	{"h", "A", "t"},
// 	{"0", "o", "h", "e", "h", "A", "t", "A"},
// }

// var Old_farm = internal.Farm{
// 	AntAmount: "9",
// 	Start:     "start 0 3",
// 	End:       "end 10 1",
// 	Links: []string{"G0-G1", "G1-G2", "G2-G3", "G3-G4", "G4-D3", "start-A0", "A0-A1", "A0-D1", "A1-A2",
// 		"A1-B1", "A2-end", "A2-C3", "start-B0", "B0-B1", "B1-E2", "start-C0", "C0-C1", "C1-C2",
// 		"C2-C3", "C3-I4", "D1-D2", "D1-F2", "D2-E2", "D2-D3", "D2-F3", "D3-end", "F2-F3", "F3-F4",
// 		"F4-end", "I4-I5", "I5-end"},
// 	Rooms: []string{"C0 1 0", "C1 2 0", "C2 3 0", "C3 4 0", "I4 5 0", "I5 6 0", "A0 1 2", "A1 2 1",
// 		"A2 4 1", "B0 1 4", "B1 2 4", "E2 6 4", "D1 6 3", "D2 7 3", "D3 8 3", "H4 4 2", "H3 5 2",
// 		"F2 6 2", "F3 7 2", "F4 8 2", "G0 1 5", "G1 2 5", "G2 3 5", "G3 4 5", "G4 6 5"},
// }
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
