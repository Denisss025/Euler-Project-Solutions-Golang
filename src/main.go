package main

import (
    "euler"
    _ "euler/solutions"
    "fmt"
)

func main() {
    euler.RunSolutions(func(name string, solution euler.SolutionFunc) {
        if result, err := solution(); err != nil {
            fmt.Println("Problem", name, "finished with error:", err)
        } else {
            fmt.Println("Problem", name + ":", result)
        }
    })
}