package main

import "goflow/flow"

func main() {
    board := flow.NewRectangleBoard(9,9)
    //board.PrintEdges()
    board.AddPathEnds(2,62)
    board.AddPathEnds(10,39)
    board.AddPathEnds(11,56)
    board.AddPathEnds(12,53)
    board.AddPathEnds(13,43)
    board.AddPathEnds(21,70)
    board.AddPathEnds(22,52)
    board.AddPathEnds(40,69)
    
    
    board.Solve()
}