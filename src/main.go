package main

import "flow"

func main() {
    board := flow.NewRectangleBoard(4,4)
    //board.PrintEdges()
    board.AddPathEnds(0,15)
    board.AddPathEnds(3,6)
    board.Solve()
}