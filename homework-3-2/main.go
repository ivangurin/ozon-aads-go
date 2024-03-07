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

	mapa := NewMapa()

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {

		var command, key, value string
		fmt.Fscan(in, &command)

		switch command {
		case "put":
			fmt.Fscan(in, &key, &value)
			mapa.Put(key, value)

		case "get":
			fmt.Fscan(in, &key)
			value, exists := mapa.Get(key)
			if exists {
				fmt.Fprintf(out, "+%s\n", value)
			} else {
				fmt.Fprintf(out, "-\n")
			}

		case "print":
			mapa.Print(out)

		}

	}

	fmt.Fscanln(in)

}

type BucketItem struct {
	Key   string
	Value string
}

type Bucket []*BucketItem

type Mapa struct {
	hashier Hashier
	buckets []Bucket
}

func NewMapa() *Mapa {

	m := &Mapa{
		hashier: *NewHashier(31),
	}

	m.init(4)

	return m

}

func (m *Mapa) Put(key, value string) {

	backet := m.getBucket(key)

	for _, item := range m.buckets[backet] {
		if item.Key == key {
			item.Value = value
			return
		}
	}

	m.buckets[backet] = append(m.buckets[backet], &BucketItem{
		Key:   key,
		Value: value,
	})

	if float64(m.getSize())/float64(m.getCapacity()) >= 0.75 {
		m.rebuild()
	}

}

func (m *Mapa) Get(key string) (string, bool) {

	backet := m.getBucket(key)

	for _, item := range m.buckets[backet] {
		if item.Key == key {
			return item.Value, true
		}
	}

	return "", false

}

func (m *Mapa) Print(out *bufio.Writer) {

	fmt.Fprintln(out, m.getSize(), m.getCapacity())

	for _, bucket := range m.buckets {

		if len(bucket) == 0 {
			fmt.Fprintln(out)
		} else {
			for _, item := range bucket {
				fmt.Fprintf(out, "\t%s %s", item.Key, item.Value)
			}
			fmt.Fprintf(out, "\n")
		}

	}

}

func (m *Mapa) init(c int) {
	m.buckets = make([]Bucket, c)

	for i := 0; i < len(m.buckets); i++ {
		m.buckets[i] = Bucket{}
	}
}

func (m *Mapa) getBucket(key string) int {
	hash := m.hashier.GetHash(key)
	return int(hash % uint64(m.getCapacity()))
}

func (m *Mapa) getSize() int {
	var res int
	for _, backet := range m.buckets {
		res += len(backet)
	}
	return res
}

func (m *Mapa) getCapacity() int {
	return len(m.buckets)
}

func (m *Mapa) rebuild() {
	buckets := m.buckets

	m.init(len(m.buckets) * 2)

	for i := 0; i < len(buckets); i++ {
		for j := len(buckets[i]) - 1; j >= 0; j-- {
			m.Put(buckets[i][j].Key, buckets[i][j].Value)
		}
	}
}

type Hashier struct {
	m uint64
}

func NewHashier(m uint64) *Hashier {
	return &Hashier{
		m: m,
	}
}

func (h *Hashier) GetHash(v string) uint64 {
	return h.hashString(v)
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
