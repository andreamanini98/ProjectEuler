package main

import "fmt"

type TreeP struct {
	l       *TreeP
	r       *TreeP
	val     int
	maxPath int
}

var mp int = 0

func main() {
	var a [][]int = make([][]int, 0)
	initializeTriangle(&a)

	var t TreeP
	buildTreeP(&a, &t)

	computeMaxPath(&t)
	fmt.Println(mp)
}

func traverseTreeP(t *TreeP) {
	if t.l != nil {
		traverseTreeP(t.l)
	}
	fmt.Print(t.val, " ")
	if t.r != nil {
		traverseTreeP(t.r)
	}
}

func computeMaxPath(t *TreeP) {
	if t.l != nil {
		newMax := t.maxPath + t.l.val
		t.l.maxPath = newMax
		if mp < newMax {
			mp = newMax
		}
		computeMaxPath(t.l)
	}
	if t.r != nil {
		newMax := t.maxPath + t.r.val
		t.r.maxPath = newMax
		if mp < newMax {
			mp = newMax
		}
		computeMaxPath(t.r)
	}
}

func buildTreeP(a *[][]int, t *TreeP) {
	var nodes [][]TreeP = getTriangleNodes(a)

	for i := len(nodes) - 2; i >= 0; i-- {
		previousNodes := nodes[i+1]
		for j := range nodes[i] {
			nodes[i][j].l = &previousNodes[j]
			nodes[i][j].r = &previousNodes[j+1]
		}
	}
	nodes[0][0].maxPath = nodes[0][0].val
	*t = nodes[0][0]
}

func getTriangleNodes(a *[][]int) [][]TreeP {
	var nodes [][]TreeP = make([][]TreeP, 0)

	for _, u := range *a {
		var nodeRow []TreeP
		for _, v := range u {
			nodeRow = append(nodeRow, TreeP{val: v})
		}
		nodes = append(nodes, nodeRow)
	}
	return nodes
}

func printTriangle(a *[][]int) {
	for _, u := range *a {
		for _, v := range u {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func initializeTriangle(a *[][]int) {
	*a = append(*a, []int{75})
	*a = append(*a, []int{95, 64})
	*a = append(*a, []int{17, 47, 82})
	*a = append(*a, []int{18, 35, 87, 10})
	*a = append(*a, []int{20, 4, 82, 47, 65})
	*a = append(*a, []int{19, 1, 23, 75, 3, 34})
	*a = append(*a, []int{88, 2, 77, 73, 7, 63, 67})
	*a = append(*a, []int{99, 65, 4, 28, 6, 16, 70, 92})
	*a = append(*a, []int{41, 41, 26, 56, 83, 40, 80, 70, 33})
	*a = append(*a, []int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29})
	*a = append(*a, []int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14})
	*a = append(*a, []int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57})
	*a = append(*a, []int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48})
	*a = append(*a, []int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31})
	*a = append(*a, []int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23})
}
