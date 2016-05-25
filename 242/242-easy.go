package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
}

func plant(people, plants int) int {
	var weeks, fruits int
	for fruits < people {
		weeks++
		fruits += plants
		plants += fruits
	}
	return weeks + 1
}

func main() {
	args := os.Args[1:]
	people, err := strconv.Atoi(args[0])
	check(err)
	plants, err := strconv.Atoi(args[1])
	check(err)
	weeks := plant(people, plants)
	fmt.Println(weeks)
}
