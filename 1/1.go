package main

import "fmt"

func main() {
	var name string
	var age string
	var username string

	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter your age")
	fmt.Scanln(&age)
	fmt.Println("Enter your username")
	fmt.Scanln(&username)
	fmt.Printf("Your name is %s, you are %s years old and your username is %s\n", name, age, username)
}
