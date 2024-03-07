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

	var a, b, c int
	fmt.Fscanln(in, &a, &b, &c)

	var n int
	fmt.Fscanln(in, &n)

	graph := map[int][]int{}
	for i := 0; i < n; i++ {
		var s, e int
		fmt.Fscanln(in, &s, &e)
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}
	fmt.Fscanln(in)

	fmt.Fprintln(out, CalcFields(a, b, c, graph))

}

func CalcFields(a, b, c int, graph map[int][]int) int {

	peaks := BFS(a, graph)
	ab := GetLength(a, b, c, peaks)
	if ab < 0 {
		return -1
	}

	peaks = BFS(b, graph)
	bc := GetLength(b, c, -1, peaks)
	if bc < 0 {
		return -1
	}

	return 1 + ab + bc

}

func BFS(start int, graph map[int][]int) map[int]int {

	queue := []int{}
	seen := map[int]int{}

	queue = append(queue, start)
	seen[start] = 0

	for i := 0; i < len(queue); i++ {

		v := queue[i]

		for _, n := range graph[v] {

			_, exists := seen[n]
			if exists {
				continue
			}

			queue = append(queue, n)
			seen[n] = v

		}

	}

	return seen

}

func GetLength(from, to, forbid int, peaks map[int]int) int {
	res := 0
	for {
		peak, exists := peaks[to]
		if !exists {
			return -1
		}

		if forbid > 0 && forbid == peak {
			return -1
		}

		res++

		if peak == from {
			break
		}

		to = peak
	}
	return res
}
