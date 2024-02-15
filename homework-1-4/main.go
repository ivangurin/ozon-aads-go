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

	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m[i])
	}
	fmt.Fscanln(in)

	fmt.Fprintln(out, countInversions(m))

}

func countInversions(m []int) int {

	res := 0
	for i := 0; i < len(m)-1; i++ {
		for j := i + 1; j < len(m); j++ {
			if m[i] > m[j] {
				res++
			}
		}
	}

	return res

}
