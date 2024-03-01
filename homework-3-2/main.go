package main

import (
	"bufio"
	"fmt"
	"math"
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

	buckets := make([]Bucket, 4)

	for i := 0; i < len(buckets); i++ {
		buckets[i] = Bucket{}
	}

	return &Mapa{
		hashier: *NewHashier(33),
		buckets: buckets,
	}

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

func (m *Mapa) getBucket(key string) int {
	hash := m.hashier.GetHash(key)
	return int(hash % uint(m.getCapacity()))
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

	m.buckets = make([]Bucket, len(buckets)*2)

	for i := 0; i < len(m.buckets); i++ {
		m.buckets[i] = Bucket{}
	}

	for _, bucket := range buckets {
		for _, item := range bucket {
			m.Put(item.Key, item.Value)
		}
	}

}

type Hashier struct {
	m uint
}

func NewHashier(m uint) *Hashier {
	return &Hashier{
		m: m,
	}
}

func (h *Hashier) GetHash(v string) uint {
	return h.hashString(v)
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
