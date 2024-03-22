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

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(in, &v)
		numbers[i] = v
	}

	fmt.Fprintln(out, maxMoney(numbers))
}

func maxMoney(boxes []int) int {
	n := len(boxes)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = boxes[0]

	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+boxes[i-1])
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
