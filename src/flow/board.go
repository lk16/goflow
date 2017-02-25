package flow

import "fmt"

type end struct {
    node_id [2]int
}

type Board struct {
    ends []end
    edges [][]int
}

func NewSquareBoard(width,height int) *Board {
    board := Board{make([]end,0),make([][]int,width*height)}
    for i:=0; i<width*height; i++ {
        board.edges[i] = make([]int,0)
    }
    for y:=0; y<height; y++ {
        for x:=0; x<width; x++ {
            field := y*width + x
            if x != 0 {
                left := y*width + x-1
                board.addEdge(field,left)
            }
            if y != 0 {
                top := (y-1)*width + x
                board.addEdge(field,top)
            }
        }
    }
    return &board

}

func (board *Board) addEdge(a,b int) {
    board.edges[a] = append(board.edges[a],b)
    board.edges[b] = append(board.edges[b],a)
}

func (board *Board) Print() {
    for id,node := range board.edges {
        fmt.Printf("Node (%d): ",id)
        for _,neighbour := range node {
            fmt.Printf("%d ",neighbour)
        }
        fmt.Printf("\n")
    }
} 