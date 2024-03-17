package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	ID       int
	Priority int
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
	heap := NewHeap()
	for i := 0; i < n; i++ {
		var x, p int
		fmt.Fscan(in, &x)
		if x == -1 {
			fmt.Fscanln(in)
			heap.Extract()
		} else {
			fmt.Fscanln(in, &p)
			heap.Insert(&HeapElement{uint64(100000000000000 + p*10000000 + x), task{x, p}})
		}
	}

	result := []int{}
	for {
		element := heap.Extract()
		if element == nil {
			break
		}
		task := element.Content.(task)
		result = append(result, task.ID)
	}

	fmt.Fprintln(out, strings.Trim(fmt.Sprint(result), "[]"))
}

type heap struct {
	array []*HeapElement
}

type HeapElement struct {
	Value   uint64
	Content interface{}
}

func NewHeap() *heap {
	array := make([]*HeapElement, 0)
	return &heap{
		array: array,
	}
}

func (h *heap) Insert(v *HeapElement) {
	h.array = append(h.array, v)
	h.siftUp(len(h.array) - 1)
}

func (h *heap) Extract() *HeapElement {
	if len(h.array) == 0 {
		return nil
	}
	res := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.siftDown(0)
	return res
}

func (h *heap) GetMin() *HeapElement {
	if len(h.array) > 0 {
		return h.array[0]
	}
	return nil
}

func (h *heap) siftUp(index int) {
	if index == 0 {
		return
	}
	if h.array[index].Value < h.array[h.getParent(index)].Value {
		h.array[index], h.array[h.getParent(index)] = h.array[h.getParent(index)], h.array[index]
		h.siftUp(h.getParent(index))
	}
}

func (h *heap) siftDown(index int) {
	if index > len(h.array)-1 {
		return
	}
	if h.getRightChild(index) <= len(h.array)-1 {
		if h.array[h.getLeftChild(index)].Value <= h.array[h.getRightChild(index)].Value {
			if h.array[h.getLeftChild(index)].Value < h.array[index].Value {
				h.array[index], h.array[h.getLeftChild(index)] = h.array[h.getLeftChild(index)], h.array[index]
				h.siftDown(h.getLeftChild(index))
			}
		} else {
			if h.array[h.getRightChild(index)].Value < h.array[index].Value {
				h.array[index], h.array[h.getRightChild(index)] = h.array[h.getRightChild(index)], h.array[index]
				h.siftDown(h.getRightChild(index))
			}
		}
	} else if h.getLeftChild(index) <= len(h.array)-1 {
		if h.array[h.getLeftChild(index)].Value < h.array[index].Value {
			h.array[index], h.array[h.getLeftChild(index)] = h.array[h.getLeftChild(index)], h.array[index]
			h.siftDown(h.getLeftChild(index))
		}
	}
}

func (h *heap) getParent(index int) int {
	return (index - 1) / 2
}

func (h *heap) getLeftChild(index int) int {
	return index*2 + 1
}

func (h *heap) getRightChild(index int) int {
	return index*2 + 2
}

func (h *heap) Print() {
	for i := 0; i < len(h.array); i++ {
		fmt.Printf("%d: array: %+v\n", i, h.array[i])
	}
}
