package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Fprintln(out, strings.Trim(fmt.Sprint(inOrder(tree)), "[]"))
	fmt.Fprintln(out, strings.Trim(fmt.Sprint(preOrder(tree)), "[]"))
	fmt.Fprintln(out, strings.Trim(fmt.Sprint(postOrder(tree)), "[]"))

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

func inOrder(node *Node) []int {
	res := []int{}

	if node.Left != nil {
		res = append(res, inOrder(node.Left)...)
	}

	res = append(res, node.Key)

	if node.Right != nil {
		res = append(res, inOrder(node.Right)...)
	}

	return res
}

func preOrder(node *Node) []int {
	res := []int{}

	res = append(res, node.Key)

	if node.Left != nil {
		res = append(res, preOrder(node.Left)...)
	}

	if node.Right != nil {
		res = append(res, preOrder(node.Right)...)
	}

	return res
}

func postOrder(node *Node) []int {
	res := []int{}

	if node.Left != nil {
		res = append(res, postOrder(node.Left)...)
	}

	if node.Right != nil {
		res = append(res, postOrder(node.Right)...)
	}

	res = append(res, node.Key)

	return res
}
