package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var n int
	fmt.Fscanln(in, &n)

	tree := make(map[int][]int)
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		tree[p] = append(tree[p], i)
	}
	fmt.Fscanln(in)

	fmt.Fprintln(out, getTreeHeight(tree, tree[-1][0]))

}

func getTreeHeight(tree map[int][]int, parent int) int {

	children, exists := tree[parent]
	if !exists {
		return 0
	}

	maxHeight := 0
	for _, child := range children {
		height := getTreeHeight(tree, child)
		if maxHeight < height {
			maxHeight = height
		}
	}

	return maxHeight + 1

}
