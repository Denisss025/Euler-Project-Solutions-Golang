package main

import (
	"euler"
	_ "euler/solutions"
	"fmt"
)

const title = `#Project Euler Answers

##Solved by

Denis Novikov

##URL

https://github.com/Denisss025/Euler-Project-Solutions-Golang

##Answers

| Problem | Result |
|:-------:| ------ |`

func main() {
	fmt.Print(title)

	fmt.Println()
	euler.RunSolutions(func(name string, solution euler.SolutionFunc) {
		fmt.Print("|   ", name, "   | ")
		if result, err := solution(); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(result)
		}
		fmt.Println(" |")
	})
}
