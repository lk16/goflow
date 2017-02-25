package flow

import "fmt"

type node struct {
    neigbours *[]node
    endpoint bool
    x,y,colour int
}

type Board struct {
    width,height int
    fields [][]node
}

func NewSquareBoard(width,height int) *Board {
    board := Board{width,height,make([][]node,height)}
    for i:=0; i<height; i++ {
        board.fields[i] = make([]node,width)
    }
    return &board

}

func (board *Board) getField(x,y int) node {
    if x < 0 || x >= board.width {
        panic("Accessing x outside bounds")
    } 
    if y < 0 || y >= board.height {
        panic("Accessing y outside bounds")
    }
    return board.fields[y][x]
}

func (board *Board) Print() {
    for y:=0; y<board.height; y++ {
        for x:=0; x<board.width; x++ {
            v := board.getField(x,y).colour
            if v == 0 {
                fmt.Printf(" -")
            } else {
                fmt.Printf(" %d",v)
            }
        }
        fmt.Printf("\n")
    }
} 