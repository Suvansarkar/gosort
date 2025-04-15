package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Distrubuted Sorting Algorithm implementation gala!")
	fmt.Println("Created by: Suvan Sarkar")
	fmt.Println("Github: https://github.com/Suvansarkar/gosort")
	fmt.Println("")
	fmt.Println("To run Odd Even transposition sort, run:")
	fmt.Println("go run oddeven/oddeven.go <number of processes>")
	fmt.Println("if number of processes are not given a default of 10, 20 and 30 will be taken")
	fmt.Println("")
	fmt.Println("To run Sasaki's Time Optimal Algorithm, run:")
	fmt.Println("go run sasaki/sasaki.go <number of processes>")
	fmt.Println("if number of processes are not given a default of 10, 20 and 30 will be taken")
	fmt.Println("")
	fmt.Println("To run Alternative Time Optimal Algorithm, run:")
	fmt.Println("go run alternative/alternative.go <number of processes>")
	fmt.Println("if number of processes are not given a default of 10, 20 and 30 will be taken")
}
