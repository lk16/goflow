package flow

import "fmt"

type pair [2]int

type Board interface {
    PrintEdges()
    Print()
    Solve()
}

type RectangleBoard struct {
    ends []pair
    edges [][]int
    width,height int
}

func NewRectangleBoard(width,height int) *RectangleBoard {
    board := RectangleBoard{make([]pair,0),make([][]int,width*height),width,height}
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

func (board *RectangleBoard) AddEnds(a,b int) {
    if a < 0 || a >= board.width*board.height {
        panic("Edge index invalid")
    }
    if b < 0 || b >= board.width*board.height {
        panic("Edge index invalid")
    }
    board.ends = append(board.ends,pair{a,b})
}

func (board *RectangleBoard) addEdge(a,b int) {
    board.edges[a] = append(board.edges[a],b)
    board.edges[b] = append(board.edges[b],a)
}

func (board *RectangleBoard) PrintEdges() {
    for id,node := range board.edges {
        fmt.Printf("Node (%d): ",id)
        for _,neighbour := range node {
            fmt.Printf("%d ",neighbour)
        }
        fmt.Printf("\n")
    }
} 

func (board *RectangleBoard) Print() {
    board.print(make([]pair,0))
}

func (board *RectangleBoard) print(edges []pair) {

    endpoints := make([]int,board.width*board.height)

    for i,_ := range endpoints {
        endpoints[i] = -1
    }

    for id,pair := range board.ends {
        endpoints[pair[0]] = id
        endpoints[pair[1]] = id
    }

    for y:=0; y<board.height; y++ {
        for x:=0; x<board.width; x++ {
            
            has_edge := false
            id := y*board.width + x
            for _,edge := range edges {
                if (edge[0] == id-1) && (edge[1] == id) {
                    has_edge = true
                    break
                }  
            } 
            if has_edge {
                fmt.Printf("---")                
            } else {
                fmt.Printf("   ")
            }

            endpoint_id := endpoints[y*board.width + x]
            if endpoint_id >= 0 {
                fmt.Printf("%d",endpoint_id)
            } else {
                fmt.Printf("·")
            }
        }
        fmt.Printf("\n")

        for x:=0; x<board.width; x++ {
            has_edge := false
            id := y*board.width + x
            for _,edge := range edges {
                if (edge[0] == id) && (edge[1] == id+board.width) {
                    has_edge = true
                    break
                }  
            } 
            if has_edge {
                fmt.Printf("---")
            } else {
                fmt.Printf("   ")
            }

        }
        fmt.Printf("\n")
    }
}

type solver struct {
    board *RectangleBoard
    edges []pair // smallest id first
    node_used []bool
}

func (board *RectangleBoard) Solve() {
    state := solver{
        board: board,
        edges: make([]pair,0),
        node_used: make([]bool,board.width*board.height)}
    state.solve()
}

func (state *solver) print() {

}

func (state *solver) solve() {

}