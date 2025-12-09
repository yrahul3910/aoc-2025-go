package day08

import (
	"fmt"
	"yrahul3910/aoc-2025-go/utils"
)

type Node[T any] struct {
	Value T

	parent *Node[T]
	size   int // for union-by-size
}

func (n Node[T]) String() string {
	return fmt.Sprint(n.Value)
}

type UnionFind[T comparable] struct {
	nodes map[T]*Node[T] // lookup by value
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{nodes: make(map[T]*Node[T])}
}

func (uf *UnionFind[T]) MakeSet(val T) {
	if _, ok := uf.nodes[val]; ok {
		return
	}

	node := &Node[T]{Value: val, size: 1}
	node.parent = node

	uf.nodes[val] = node
}

func (uf *UnionFind[T]) Find(node *Node[T]) *Node[T] {
	// path halving
	for node.parent != node {
		node.parent = node.parent.parent
		node = node.parent
	}

	return node
}

func (uf *UnionFind[T]) Union(x, y *Node[T]) {
	// replace nodes by roots
	x = uf.Find(x)
	y = uf.Find(y)

	if x == y {
		return
	}

	if x.size < y.size {
		z := x
		x = y
		y = z
	}

	y.parent = x
	x.size += y.size
}

func (uf *UnionFind[T]) FindValue(val T) *Node[T] {
	node, ok := uf.nodes[val]
	if !ok {
		return nil
	}

	return uf.Find(node)
}

func (uf *UnionFind[T]) UnionValues(a, b T) {
	nodeA, okA := uf.nodes[a]
	nodeB, okB := uf.nodes[b]
	if !okA || !okB {
		return
	}

	uf.Union(nodeA, nodeB)
}

type Component[T comparable] struct {
	Value T
	Count int
}

func (uf *UnionFind[T]) PrintComponent(c Component[T]) {
	nodes := make([]Node[T], 0)

	for _, node := range uf.nodes {
		if root := uf.Find(node); root.Value == c.Value {
			nodes = append(nodes, *node)
		}
	}

	fmt.Print("Component {\n\tnodes = ")
	utils.PrintArray(nodes)
	fmt.Printf("\n\tsize = %d\n}", c.Count)
}

func (uf *UnionFind[T]) Components() []*Component[T] {
	indexSet := map[*Node[T]]int{}
	components := make([]*Component[T], 0)

	for _, node := range uf.nodes {
		root := uf.Find(node)

		index, ok := indexSet[root]
		if ok {
			components[index].Count++
		} else {
			newComponent := &Component[T]{root.Value, 1}
			components = append(components, newComponent)
			indexSet[root] = len(components) - 1
		}
	}

	return components
}
