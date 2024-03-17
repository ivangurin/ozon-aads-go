package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type attrib struct {
	ID    int
	Value int
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var m, n, e, t int
	fmt.Fscanln(in, &m, &n, &e, &t)

	attribCounter := 0
	attribs := map[attrib]int{}
	attribsIDs := []int{}
	clients := map[int]map[int]struct{}{}
	graph := NewGraph(m)
	for i := 0; i < e; i++ {
		var c, a, v int
		fmt.Fscanln(in, &c, &a, &v)

		graph.AddVertex(c)

		attrib := attrib{a, v}
		if _, exists := attribs[attrib]; !exists {
			attribCounter++
			attribsIDs = append(attribsIDs, attribCounter)
			attribs[attrib] = attribCounter
		}

		if _, exists := clients[c]; !exists {
			clients[c] = map[int]struct{}{}
		}

		clients[c][attribs[attrib]] = struct{}{}

	}
	fmt.Fscanln(in)

	attribCombs := GetCombs(attribsIDs, t, 0)

	for _, attribComb := range attribCombs {

		clientsGroup := []int{}
		for clientID, client := range clients {

			suitable := true
			for _, attribID := range attribComb {
				if _, exists := client[attribID]; !exists {
					suitable = false
					break
				}
			}

			if suitable {
				clientsGroup = append(clientsGroup, clientID)
			}
		}

		if len(clientsGroup) > 1 {
			for _, clientFrom := range clientsGroup {
				for _, clientTo := range clientsGroup {
					if clientFrom == clientTo {
						continue
					}
					graph.AddEdge(clientFrom, clientTo)
				}
			}
		}
	}

	opgs := graph.FindGroups()
	if len(opgs) > 0 {
		fmt.Fprintln(out, strings.Trim(fmt.Sprint(opgs[0]), "[]"))
	}

	// fmt.Println("attribs: ", attribs)
	// fmt.Println("attribsIDs: ", attribsIDs)
	// fmt.Println("attribCombs: ", attribCombs)
	// fmt.Println("clients: ", clients)
	// graph.Print()
	// fmt.Println(opgs)

}

func GetCombs(list []int, numbers int, from int) [][]int {
	res := [][]int{}
	for i, number := range list[from:] {
		if numbers > 1 {
			subres := GetCombs(list, numbers-1, from+i+1)
			for _, subNumbers := range subres {
				resNumbers := append([]int{number}, subNumbers...)
				res = append(res, resNumbers)
			}
		} else {
			res = append(res, []int{number})
		}
	}
	return res
}

type Vertex struct {
	Edges map[int]struct{}
}

type Graph struct {
	vertexes map[int]*Vertex
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertexes: make(map[int]*Vertex, v),
	}
}

func (g *Graph) AddVertex(id int) *Vertex {
	if _, exists := g.vertexes[id]; !exists {
		g.vertexes[id] = &Vertex{
			Edges: map[int]struct{}{},
		}
	}
	return g.vertexes[id]
}

func (g *Graph) AddEdge(fromVertexID int, toVertexID int) {
	vertex := g.AddVertex(fromVertexID)
	vertex.Edges[toVertexID] = struct{}{}
	g.AddVertex(toVertexID)
}

func (g *Graph) FindGroups() [][]int {
	groups := []map[int]struct{}{}
	for vertexID := range g.vertexes {

		ingroup := false
		for _, group := range groups {
			if _, exists := group[vertexID]; exists {
				ingroup = true
				break
			}
		}
		if ingroup {
			continue
		}

		group := map[int]struct{}{}
		g.DFS(vertexID, group)

		groups = append(groups, group)
	}

	res := make([][]int, 0, len(groups))
	for _, group := range groups {
		resGroup := make([]int, 0, len(group))
		for vertextID := range group {
			resGroup = append(resGroup, vertextID)
		}
		sort.Ints(resGroup)
		res = append(res, resGroup)
	}

	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) > len(res[j]) {
			return true
		}
		if len(res[i]) == len(res[j]) {
			return res[i][0] < res[j][0]
		}
		return false
	})

	return res
}

func (g *Graph) DFS(vertexID int, seen map[int]struct{}) {
	if _, exists := seen[vertexID]; exists {
		return
	}
	seen[vertexID] = struct{}{}
	vertex := g.vertexes[vertexID]
	for toVertexID := range vertex.Edges {
		g.DFS(toVertexID, seen)
	}
}

func (g *Graph) Print() {
	for vertexID, vertex := range g.vertexes {
		fmt.Printf("%d: %+v\n", vertexID, vertex)
	}
}
