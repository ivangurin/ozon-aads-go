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
		var command, key string
		fmt.Fscan(in, &command)

		switch command {
		case "add":
			fmt.Fscan(in, &key)
			mapa.Add(key)

		case "remove":
			fmt.Fscan(in, &key)
			mapa.Remove(key)

		case "contains":
			fmt.Fscan(in, &key)
			exists := mapa.Contains(key)
			if exists {
				fmt.Fprint(out, "+\n")
			} else {
				fmt.Fprint(out, "-\n")
			}

		case "print":
			mapa.Print(out)

		}
	}
}

type BucketItem struct {
	Key string
}

type Bucket []*BucketItem

type Mapa struct {
	buckets []Bucket
}

func NewMapa() *Mapa {
	m := &Mapa{}
	m.init(4)
	return m
}

func (m *Mapa) Add(key string) {
	bucketNumber := m.getBucket(key)

	for _, item := range m.buckets[bucketNumber] {
		if item.Key == key {
			return
		}
	}

	m.buckets[bucketNumber] = append(m.buckets[bucketNumber], &BucketItem{
		Key: key,
	})

	if float64(m.getSize())/float64(m.getCapacity()) >= 0.75 {
		m.rebuild()
	}
}

func (m *Mapa) Contains(key string) bool {
	bucketNumber := m.getBucket(key)

	for _, item := range m.buckets[bucketNumber] {
		if item.Key == key {
			return true
		}
	}

	return false
}

func (m *Mapa) Remove(key string) {
	bucketNumber := m.getBucket(key)

	for ind, item := range m.buckets[bucketNumber] {
		if item.Key == key {
			bucket := Bucket{}
			bucket = append(bucket, m.buckets[bucketNumber][:ind]...)
			bucket = append(bucket, m.buckets[bucketNumber][ind+1:]...)
			m.buckets[bucketNumber] = bucket
			return
		}
	}
}

func (m *Mapa) Print(out *bufio.Writer) {
	fmt.Fprintln(out, m.getSize(), m.getCapacity())

	for _, bucket := range m.buckets {
		if len(bucket) == 0 {
			fmt.Fprintln(out)
		} else {
			for _, item := range bucket {
				fmt.Fprintf(out, " %s", item.Key)
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
	hash := Hash(key, 31)
	return int(hash % uint64(m.getCapacity()))
}

func (m *Mapa) getSize() int {
	var res int
	for _, bucket := range m.buckets {
		res += len(bucket)
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
			m.Add(buckets[i][j].Key)
		}
	}
}

func Hash(s string, m uint64) uint64 {
	var res uint64
	for i := range s {
		res = res*m + (uint64(s[i]) - 33)
	}
	return res
}
