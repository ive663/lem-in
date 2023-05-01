package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Farm struct {
	antAmount string
	start     string
	end       string
	links     []string
	rooms     []string
}

func main() {

	frm := &Farm{}

	filename := os.Args[1]
	err := frm.prepFarm(filename)
	check(err)

	fmt.Println("=================================")
	// fmt.Println("ALL: ", frm)
	fmt.Println("Start: ", frm.start)
	fmt.Println("End: ", frm.end)
	fmt.Println("antAmunt: ", frm.antAmount)
	fmt.Println("links: ", frm.links)
	fmt.Println("rooms: ", frm.rooms)

	// paths := path(links, start, end)
	// fmt.Println(paths)
}
func (frm *Farm) prepFarm(filename string) error {
	f, err := os.Open(filename)
	check(err)
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

	frm.antAmount = fileLines[0]

	for i, line := range fileLines[1:] {
		switch i {
		case idxOfStart, idxOfEnd:
			if i == idxOfStart {
				frm.start = line
			} else {
				frm.end = line
			}
		default:
			if strings.Contains(line, "-") {
				frm.links = append(frm.links, line)
			} else if !strings.Contains(line, "#") {
				frm.rooms = append(frm.rooms, line)
			}
		}
	}

	fmt.Println(fileLines)
	return nil
}

// func path(links []string, start string, end string) [][]string {
// 	p := make([][]string, 1)
// 	startloc := strings.Split(start, " ")
// 	// counter := 0
// 	p[0] = append(p[0], startloc[0])
// 	tmp := []string{}
// 	// for i1 := 0; i1 <= counter; i1++ {
// 	for i, val := range links {
// 		if strings.Contains(val, p[i][i]) {

// 			tmp = append(tmp, links[i])

func remove(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

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

func check(e error) {
	if e != nil {
		log.Println("err:", e)
		return

	}
}
