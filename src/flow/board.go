package flow

import "fmt"

type pair [2]int

type RectangleBoard struct {
    path_ends []pair
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

func (board *RectangleBoard) AddPathEnds(a,b int) {
    if a < 0 || a >= board.width*board.height {
        panic("Edge index invalid")
    }
    if b < 0 || b >= board.width*board.height {
        panic("Edge index invalid")
    }
    board.path_ends = append(board.path_ends,pair{a,b})
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

    for id,pair := range board.path_ends {
        endpoints[pair[0]] = id
        endpoints[pair[1]] = id
    }

    for y:=0; y<board.height; y++ {
        for x:=0; x<board.width; x++ {
            
            has_edge := false
            id := y*board.width + x
            for _,edge := range edges {
                if ((edge[0] == id-1) && (edge[1] == id)) || ((edge[1] == id-1) && (edge[0] == id)) {
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
                if ((edge[0] == id) && (edge[1] == id+board.width)) || ((edge[1] == id) && (edge[0] == id+board.width)) {
                    has_edge = true
                    break
                }  
            } 
            if has_edge {
                fmt.Printf("   |")
            } else {
                fmt.Printf("    ")
            }

        }
        fmt.Printf("\n")
    }

    fmt.Printf("\n\n\n")
}

type Solver struct {
    board *RectangleBoard
    edges []pair
    node_used []bool
    node_used_count int
}

func (board *RectangleBoard) Solve() {
    solver := Solver{
        board: board,
        edges: make([]pair,0),
        node_used: make([]bool,board.width*board.height),
        node_used_count: 0}
    for _,p := range board.path_ends {
        solver.node_used[p[0]] = true
        solver.node_used[p[1]] = true
        solver.node_used_count += 2
    }
    solver.solvePath(0)
}

func (solver *Solver) print() {
    solver.board.print(solver.edges)
}

func (solver *Solver) solvePath(path_id int) bool {

    if path_id == len(solver.board.path_ends) {
        return solver.node_used_count == len(solver.board.edges)           
    }

    ends := solver.board.path_ends[path_id]

    // path endpoint is unused for this path
    solver.node_used[ends[1]] = false
    solver.node_used_count -= 1

    result := solver.solvePathRec(ends[0],ends[1],path_id)

    // pretend this node IS used for other paths
    solver.node_used[ends[1]] = true
    solver.node_used_count += 1

    return result
}

func (solver *Solver) solvePathRec(from,to,path_id int) bool {
    
    if from == to {
        if solver.solvePath(path_id+1) {
            return true
        }
    }

    neighbours := solver.board.edges[from]

    for _,next := range neighbours {
        if !solver.node_used[next] {
            solver.node_used[next] = true
            solver.node_used_count += 1
            solver.edges = append(solver.edges,pair{from,next})

            fmt.Printf("nodes_used (%d): ",solver.node_used_count)
            for i,_ := range solver.node_used {
                if solver.node_used[i] {
                    fmt.Printf("%d ",i)
                }
            }
            fmt.Printf("\n")
            solver.print()
         
            if solver.solvePathRec(next,to,path_id) {
                return true
            }

            solver.edges = solver.edges[:len(solver.edges)-1]
            solver.node_used[next] = false
            solver.node_used_count -= 1
            
        }
    }

    return false
}
