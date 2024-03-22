package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n1 int
	fmt.Fscanln(in, &n1)
	word1 := make([]string, n1)
	for i := 0; i < n1; i++ {
		v, _ := in.ReadString('\n')
		word1[i] = strings.TrimSuffix(v, "\n")
	}

	var n2 int
	fmt.Fscanln(in, &n2)
	word2 := make([]string, n2)
	for i := 0; i < n2; i++ {
		v, _ := in.ReadString('\n')
		word2[i] = strings.TrimSuffix(v, "\n")
	}

	fmt.Fprintln(out, CalcLevDist(word1, word2))
}

func CalcLevDist(word1, word2 []string) int {

	if debug {
		fmt.Println(word1)
		fmt.Println(word2)
	}

	prevRow := make([]int, len(word2)+1)
	currRow := make([]int, len(word2)+1)
	for i := 0; i < len(prevRow); i++ {
		prevRow[i] = i
	}

	if debug {
		fmt.Println("   ", word2)
		fmt.Println(" ", prevRow)
	}

	for row := 1; row <= len(word1); row++ {
		currRow[0] = row
		for col := 1; col <= len(word2); col++ {
			currRow[col] = min(currRow[col-1]+1, prevRow[col]+1, prevRow[col-1]+m(word1[row-1], word2[col-1]))
		}
		if debug {
			fmt.Println(word1[row-1], currRow)
		}
		copy(prevRow, currRow)
	}

	return currRow[len(currRow)-1]
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func m(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	return 1
}
