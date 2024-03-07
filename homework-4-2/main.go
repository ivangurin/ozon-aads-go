package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
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

	graph := map[int][]int{}
	for i := 0; i < n; i++ {
		var s, e int
		fmt.Fscanln(in, &s, &e)
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}
	fmt.Fscanln(in)

	fmt.Fprintln(out, strings.Trim(fmt.Sprint(FindCircle(graph)), "[]"))

}

func FindCircle(graph map[int][]int) []int {

	var start int
	for k := range graph {
		start = k
	}

	color := map[int]int{}
	DFS(start, -1, graph, color)

	res := []int{}
	for k, v := range color {
		if v == 2 {
			res = append(res, k)
		}
	}

	if len(res) == 0 {
		res = append(res, -1)
	} else {
		sort.Ints(res)
	}

	return res

}

func DFS(vertex int, parent int, graph map[int][]int, color map[int]int) (bool, error) {

	color[vertex] = 1

	for _, neighbour := range graph[vertex] {

		if neighbour == parent {
			continue
		}

		if color[neighbour] == 0 {

			exists, err := DFS(neighbour, vertex, graph, color)
			if err != nil {
				return false, err
			}

			if exists {
				if color[neighbour] == 1 {
					color[neighbour] = 2
					return true, nil
				} else if color[neighbour] == 2 {
					return false, errors.New("")
				}
			}

		} else if color[neighbour] == 1 {
			color[neighbour] = 2
			return true, nil
		}

	}

	return false, nil

}
