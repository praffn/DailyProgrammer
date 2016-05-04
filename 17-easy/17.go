// write an application which will print a triangle of stars of user-specified height, with each line having twice as many stars as the last. sample output:
// @
// @@
// @@@@
// hint: in many languages, the "+" sign concatenates strings.
// bonus features: print the triangle in reverse, print the triangle right justified
package main

import (
	"fmt"
	"strings"
)

func main() {
	var sym string = "@"
	var stars int
	fmt.Println("Enter amount of stars")
	fmt.Scanf("%d", &stars)

	for i := 1; i <= stars; i++ {
		line := strings.Repeat(sym, i)
		fmt.Println(line)
	}

}
