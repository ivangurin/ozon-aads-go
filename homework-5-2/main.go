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
	var v, e int
	fmt.Fscanln(in, &v, &e)
	var s, t int
	fmt.Fscanln(in, &s, &t)
	graph := NewGraph(v)
	for i := 0; i < e; i++ {
		var v1, v2, c int
		fmt.Fscanln(in, &v1, &v2, &c)
		graph.AddEdge(v1, v2, c)
	}
	fmt.Fprintln(out, graph.GetCost(s, t))
	// graph.Print()
}

type Vertex struct {
	ID           int
	Edges        map[int]int
	FromVertexID int
	Cost         int
}

type Graph struct {
	vertexes map[int]*Vertex
	heap     *Heap
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertexes: make(map[int]*Vertex, v),
		heap:     NewHeap(),
	}
}

func (g *Graph) AddVertex(id int) *Vertex {
	if _, exists := g.vertexes[id]; !exists {
		g.vertexes[id] = &Vertex{
			ID:    id,
			Edges: map[int]int{},
			Cost:  math.MaxInt64,
		}
	}
	return g.vertexes[id]
}

func (g *Graph) AddEdge(fromVertexID int, toVertexID int, cost int) {
	vertex := g.AddVertex(fromVertexID)
	vertex.Edges[toVertexID] = cost
	g.AddVertex(toVertexID)
}

func (g *Graph) GetCost(s, t int) int {
	if s == t {
		return 0
	}
	if _, exists := g.vertexes[s]; !exists {
		return -1
	}
	if _, exists := g.vertexes[t]; !exists {
		return -1
	}

	g.start(s)

	if g.vertexes[t].Cost == math.MaxInt64 {
		return -1
	} else {
		return g.vertexes[t].Cost
	}
}

func (g *Graph) start(s int) {
	g.vertexes[s].Cost = 0
	g.heap.Insert(&HeapElement{g.vertexes[s].Cost, g.vertexes[s]})

	for {
		next := g.next()
		if next == nil {
			break
		}
		g.relax(next)
	}
}

func (g *Graph) next() *Vertex {
	el := g.heap.Extract()
	if el == nil {
		return nil
	}
	return el.Content.(*Vertex)
}

func (g *Graph) relax(vertex *Vertex) {
	for toVertexID, edgeCost := range vertex.Edges {
		cost := vertex.Cost + edgeCost
		if g.vertexes[toVertexID].Cost <= cost {
			continue
		}

		g.vertexes[toVertexID].Cost = cost
		g.vertexes[toVertexID].FromVertexID = vertex.ID

		g.heap.Insert(&HeapElement{g.vertexes[toVertexID].Cost, g.vertexes[toVertexID]})
	}
}

func (g *Graph) Print() {
	for _, vertex := range g.vertexes {
		fmt.Printf("%d: %+v\n", vertex.ID, vertex)
		for edgeID, edge := range vertex.Edges {
			fmt.Printf("\t%d: %+v\n", edgeID, edge)
		}
	}
}

type Heap struct {
	array []*HeapElement
}

type HeapElement struct {
	Value   int
	Content interface{}
}

func NewHeap() *Heap {
	array := make([]*HeapElement, 0)
	return &Heap{
		array: array,
	}
}

func (h *Heap) Insert(v *HeapElement) {
	h.array = append(h.array, v)
	h.siftUp(len(h.array) - 1)
}

func (h *Heap) Extract() *HeapElement {
	if len(h.array) == 0 {
		return nil
	}
	res := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.siftDown(0)
	return res
}

func (h *Heap) siftUp(index int) {
	if index == 0 {
		return
	}
	if h.array[index].Value < h.array[h.getParent(index)].Value {
		h.array[index], h.array[h.getParent(index)] = h.array[h.getParent(index)], h.array[index]
		h.siftUp(h.getParent(index))
	}
}

func (h *Heap) siftDown(index int) {
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

func (h *Heap) getParent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) getLeftChild(index int) int {
	return index*2 + 1
}

func (h *Heap) getRightChild(index int) int {
	return index*2 + 2
}
