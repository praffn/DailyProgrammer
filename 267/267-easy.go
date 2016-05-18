package main

import (
	"fmt"
	"os"
	"strconv"
)

func ordinalize(n int) string {
	suffix := "th"
	sp := []int{11, 12, 13}

	for _, v := range sp {
		if n == v {
			return strconv.Itoa(n) + suffix
		}
	}

	switch n % 10 {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	}

	return strconv.Itoa(n) + suffix
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var places int

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(1)
	}

	place, err := strconv.Atoi(args[0])
	check(err)
	var places int

	if len(args) == 1 {
		places = 100
	} else {
		places, _ = strconv.Atoi(args[1])
	}

	for i := 1; i <= places; i++ {
		if i == place {
			continue
		}
		fmt.Println(ordinalize(i))
	}

}
