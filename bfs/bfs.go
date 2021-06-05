package bfs

import (
	"fmt"
	"strconv"
)

// TODO validate inputs

// runQuery gets a query inputs and returns different nodes distances for this query
func runQuery() []int {
	//fmt.Printf("enter nodes and edges\n")
	var nodeInput, edgesInput string
	_, _ = fmt.Scan(&edgesInput)
	_, _ = fmt.Scan(&nodeInput)
	nodesNbr, _ := strconv.Atoi(edgesInput)
	edgesNbr, _ := strconv.Atoi(nodeInput)
	//fmt.Printf("you have %v edges and %v nodes\n", edgesNbr, nodesNbr)
	edges := getEdges(edgesNbr)
	//fmt.Println("enter the starting node number")
	var inputS string
	_, _ = fmt.Scan(&inputS)
	s, _ := strconv.Atoi(inputS)
	distances := getDistances(s, edges, nodesNbr)
	// find the unlinked nodes
	for i := 1; i <= nodesNbr; i++ {
		exist := false
		for _, e := range edges {
			if e[0] == i || e[1] == i {
				exist = true
				break
			}
		}
		if !exist {
			distances = append(distances, -1)
		}
	}
	return distances
}

// getDistances gets the different distances between nodes
func getDistances(s int, edges [][2]int, n int) []int {
	mEdges := getMainEdges(edges, s)
	var distances []int
	for _, me := range mEdges {
		d := 6
		distances = append(distances, d)
		for true {
			index, ok := findInSlice(me[1], edges)
			if !ok {
				break
			}
			d = d + 6
			distances = append(distances, d)
			me = edges[index]
		}
	}

	return distances
}

// getMainEdges gets the main edges [mEdges] from the given [edges] that starts by the "starting node [s]"
func getMainEdges(edges [][2]int, s int) [][2]int {
	// define main edges "mEdges" starting by the starting node "s"
	var mEdges [][2]int
	for i, v := range edges {
		if v[0] == s {
			mEdges = append(mEdges, v)
		}
		if v[1] == s {
			v = swap(v[0], v[1])
			mEdges = append(mEdges, v)
			edges[i] = v
		}
	}
	return mEdges
}

// getEdges gets the edges from stdin and returns a slice of slices [edges]
func getEdges(e int) [][2]int {
	edges := make([][2]int, e)
	for j := 0; j < e; j++ {
		//fmt.Printf("enter the edge number %v\n", j+1)
		var uInput, vInput string
		_, _ = fmt.Scan(&uInput)
		_, _ = fmt.Scan(&vInput)
		u, _ := strconv.Atoi(uInput) //nbr of nodes
		v, _ := strconv.Atoi(vInput) // nbr of edges
		edges[j] = [2]int{u, v}
	}
	return edges
}

// findInSlice returns if a value [val] exists in a slice of slices [slice]
// if it exists the function returns true and the index of the slice
// otherwise it returns false and -1
func findInSlice(val int, slice [][2]int) (int, bool) {
	for i, s := range slice {
		if s[0] == val {
			return i, true
		}
	}
	return -1, false
}

func swap(a, b int) [2]int {
	b, a = a, b
	return [2]int{a, b}
}

/*
EXAMPLE

	// q = nbr of queries
	var q int
	//fmt.Println("ENTER THE NUMBER OF QUERIES (GRAPHS)")
	_, _ = fmt.Scan(&q)
	//fmt.Printf("you have %v queries\n", q)
	var queryDistances = make([][]int, 0, q)
	for i := 0; i < q; i++ {
		distances := runQuery()
		queryDistances = append(queryDistances, distances)
	}
	for _, d := range queryDistances {
		for _, dd := range d {
			fmt.Print(dd," ")
		}
		fmt.Println()
	}
*/
