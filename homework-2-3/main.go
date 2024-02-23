package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var n int
	fmt.Fscanln(in, &n)

	nodes := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		var key, left, right int
		fmt.Fscan(in, &key, &left, &right)
		nodes = append(nodes, []int{key, left, right})
	}
	fmt.Fscanln(in)

	tree := &Node{}
	buildTree(nodes, 0, tree)

	fmt.Fprintln(out, isSearchTree(tree))

}

func buildTree(nodes [][]int, row int, node *Node) {

	nodeData := nodes[row]

	node.Key = nodeData[0]

	if nodeData[1] != -1 {
		node.Left = &Node{}
		buildTree(nodes, nodeData[1], node.Left)
	}

	if nodeData[2] != -1 {
		node.Right = &Node{}
		buildTree(nodes, nodeData[2], node.Right)
	}

}

func isSearchTree(node *Node) string {

	if node.Left != nil {
		max := getMaxKey(node.Left)
		if max >= node.Key {
			return "no"
		}
		res := isSearchTree(node.Left)
		if res == "no" {
			return "no"
		}
	}

	if node.Right != nil {
		min := getMinKey(node.Right)
		if min < node.Key {
			return "no"
		}
		res := isSearchTree(node.Right)
		if res == "no" {
			return "no"
		}
	}

	return "yes"

}

func getMinKey(node *Node) int {
	res := node.Key
	if node.Left != nil {
		resLeft := getMinKey(node.Left)
		if resLeft < res {
			res = resLeft
		}
	}
	if node.Right != nil {
		resRight := getMinKey(node.Right)
		if resRight < res {
			res = resRight
		}
	}
	return res
}

func getMaxKey(node *Node) int {
	res := node.Key
	if node.Left != nil {
		resLeft := getMaxKey(node.Left)
		if resLeft > res {
			res = resLeft
		}
	}
	if node.Right != nil {
		resRight := getMaxKey(node.Right)
		if resRight > res {
			res = resRight
		}
	}
	return res
}
