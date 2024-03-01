package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Hashier struct {
	m uint
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

func NewHashier(m uint) *Hashier {
	return &Hashier{
		m: m,
	}
}

func (h *Hashier) GetHash(t, v string) uint {

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

func (h *Hashier) hashNumber(v string) uint {

	num, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}

	var res uint
	if num >= 0 {
		res = uint(num)
	} else {
		res = math.MaxUint64 - uint(math.Abs(float64(num))) + 1
	}

	return res

}

func (h *Hashier) hashCharacter(v string) uint {
	return getAsciiCode(v) - 33
}

func (h *Hashier) hashString(v string) uint {
	var res float64
	for i := range v {
		res += float64(h.hashCharacter(string(v[i]))) * math.Pow(float64(h.m), float64(len(v)-1-i))
	}
	return uint(res)
}

func getAsciiCode(s string) uint {
	return uint([]byte(s)[0])
}
