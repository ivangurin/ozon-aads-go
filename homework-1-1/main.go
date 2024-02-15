package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var lenM, lenN int
	fmt.Fscanln(in, &lenM, &lenN)

	m := make([]int, lenM)
	for i := 0; i < lenM; i++ {
		fmt.Fscan(in, &m[i])
	}
	fmt.Fscanln(in)

	n := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		fmt.Fscan(in, &n[i])
	}
	fmt.Fscanln(in)

	fmt.Fprintln(out, strings.Trim(fmt.Sprint(merge(m, n)), "[]"))

}

func merge(m, n []int) []int {

	res := make([]int, 0, len(m)+len(n))

	curM, curN := 0, 0
	for {

		if curM > len(m)-1 {
			res = append(res, n[curN:]...)
			break
		}

		if curN > len(n)-1 {
			res = append(res, m[curM:]...)
			break
		}

		if m[curM] <= n[curN] {
			res = append(res, m[curM])
			curM++
		} else {
			res = append(res, n[curN])
			curN++
		}

	}

	return res

}
