package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Hashier struct {
	m uint64
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var n int
	fmt.Fscanln(in, &n)

	hashier := NewHashier(127)

	for i := 0; i < n; i++ {
		var t, v string
		fmt.Fscan(in, &t, &v)
		fmt.Fprintln(out, hashier.GetHash(t, v))
	}
	fmt.Fscanln(in)

}

func NewHashier(m uint64) *Hashier {
	return &Hashier{
		m: m,
	}
}

func (h *Hashier) GetHash(t, v string) uint64 {

	switch t {
	case "number":
		return h.hashNumber(v)
	case "character":
		return h.hashCharacter(v)
	case "string":
		return h.hashString(v)
	}

	return 0

}

func (h *Hashier) hashNumber(v string) uint64 {

	num, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint64(num)

}

func (h *Hashier) hashCharacter(v string) uint64 {
	return getAsciiCode(v) - 33
}

func (h *Hashier) hashString(v string) uint64 {
	var res uint64
	for i := range v {
		res = h.m*res + h.hashCharacter(string(v[i]))
	}
	return res
}

func getAsciiCode(s string) uint64 {
	return uint64([]byte(s)[0])
}
