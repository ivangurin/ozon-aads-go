package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug = false

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	scanner := bufio.NewScanner(in)

	scanner.Scan()
	line := scanner.Text()

	var n int
	n, _ = strconv.Atoi(line)
	b := make([][]int32, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		list := strings.Split(line, " ")
		b[i] = Sign(list)
	}

	scanner.Scan()
	line = scanner.Text()

	var m int
	m, _ = strconv.Atoi(line)
	a := make([][]int32, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		list := strings.Split(line, " ")
		a[i] = Sign(list)
	}

	for _, aa := range a {
		for j, bb := range b {
			same := 0
			for i := 0; i < len(aa); i++ {
				if aa[i] == bb[i] {
					same++
				}
			}
			if j > 0 {
				if debug {
					fmt.Print(" ")
				}
				fmt.Fprint(out, " ")
			}
			if debug {
				fmt.Printf("%.3f", float64(same)/float64(len(aa)))
			}
			fmt.Fprintf(out, "%.3f", float64(same)/float64(len(aa)))
		}
		if debug {
			fmt.Println()
		}
		fmt.Fprintln(out)
	}
}

func Sign(list []string) []int32 {
	numbers := []int32{
		2, 3, 5, 7, 13, 17, 19, 31, 61, 89,
		107, 127, 521, 607, 1279, 2203, 2281, 3217, 4253, 4423,
		9689, 9941, 11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049,
		216091, 756839, 859433, 1257787, 1398269, 2976221, 3021377, 6972593, 13466917, 20996011,
		24036583, 25964951, 30402457, 32582657, 37156667, 42643801, 43112609, 57885161, 74207281, 77232917,
		82589933}

	res := make([]int32, len(numbers))
	for i, number := range numbers {
		res[i] = MinHash(list, number)
	}

	return res
}

func MinHash(list []string, number int32) int32 {
	var res int32
	for i, element := range list {
		hash := Hash(element, number)
		if i == 0 || hash < res {
			res = hash
		}
	}
	return res
}

func Hash(s string, m int32) int32 {
	var res int32
	for i := range s {
		res = res*m + (int32(s[i]) - 33)
	}
	return res * int32(m)
}
