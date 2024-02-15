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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscanln(in)

	f, l := find(a, a[len(a)-1])

	fmt.Fprintln(out, f, l)

}

func find(m []int, pivot int) (int, int) {

	var lenLeft int
	var first int = -1
	var last int = -1
	var ind int = -1
	for i := range m {
		if m[i] < pivot {
			lenLeft++
		}
		if m[i] == pivot {
			ind++
			if first < 0 {
				first = ind
			}
			last = ind
		}

	}

	if first < 0 {
		first = 0
	}

	if last < 0 {
		last = 0
	}

	return lenLeft + first, lenLeft + last

}
