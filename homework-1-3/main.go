
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

	var k int
	fmt.Fscanln(in, &k)

	for i := range a {
		ind := find(a, a[i])
		if ind == k {
			fmt.Fprintln(out, a[i])
			break
		}
	}

}

func find(m []int, v int) int {

	var lenLeft int
	for i := range m {
		if m[i] > v {
			lenLeft++
		}
	}

	return lenLeft + 1

}
