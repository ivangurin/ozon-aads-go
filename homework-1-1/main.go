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
		if i == lenM-1 {
			fmt.Fscanln(in, &m[i])
			break
		}
		fmt.Fscan(in, &m[i])
	}

	n := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		if i == lenN-1 {
			fmt.Fscanln(in, &n[i])
			break
		}
		fmt.Fscan(in, &n[i])
	}

	r := merge(m, n)

	fmt.Println(strings.Trim(fmt.Sprint(r), "[]"))

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
