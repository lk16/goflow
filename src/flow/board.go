package flow

import "fmt"

type Board struct {
    width,height int
    fields []int
}

func NewBoard(width,height int) *Board {
    board := Board{width,height,make([]int,width*height)}
    return &board
}

func (board *Board) GetField(x,y int) int {
    if x < 0 || x >= board.width {
        panic("Accessing x outside bounds")
    } 
    if y < 0 || y >= board.height {
        panic("Accessing y outside bounds")
    }
    return board.fields[y*board.width + x]
}

func (board *Board) Print() {
    for y:=0; y<board.height; y++ {
        for x:=0; x<board.width; x++ {
            v := board.GetField(x,y)
            if v == 0 {
                fmt.Printf(" -")
            } else {
                fmt.Printf(" %d",v)
            }
        }
        fmt.Printf("\n")
    }
} 