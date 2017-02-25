package main

import "flow"
import "fmt"

func main() {
    board := flow.NewRectangleBoard(4,4)
    //board.PrintEdges()
    board.AddEnds(0,15)
    fmt.Printf("\n")
    board.Print()
    fmt.Printf("\n")
    board.Solve()
}