package main

import "fmt"

type SocialGraph struct {
	nodes []*Node
}

type Node struct {
	edges     []*Node
	FirstName string
	LastName  string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.FirstName)
}

func (g *SocialGraph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *SocialGraph) AddEdge(n1, n2 *Node) {
	if n1.edges == nil {
		n1.edges = make([]*Node, 0)
	}
	if n2.edges == nil {
		n2.edges = make([]*Node, 0)
	}
	n1.edges = append(n1.edges, n2)
	n2.edges = append(n2.edges, n1)
}

func (g *SocialGraph) IsRelated(mainNode, otherNode *Node) bool {
	// need to keep track of visited nodes
	visited := make(map[*Node]struct{}) // basically a set implementation in go
	var empty struct{}                  // empty struct as a placeholder to make set works
	// keep track of a queue of nodes
	queue := []*Node{mainNode}
	for {
		node := queue[0]
		visited[node] = empty // add node to set
		queue = queue[1:]     // pop node
		queue = getNeighbors(node, visited, queue)
		if inNeighbor(queue, otherNode) {
			return true
		}
		if len(queue) == 0 {
			return false
		}
	}
}

func getNeighbors(parentNode *Node, visited map[*Node]struct{}, queue []*Node) []*Node {
	for _, n := range parentNode.edges {
		if _, ok := visited[n]; ok { // already visited
			continue
		}
		if alreadyInQueue(queue, n) {
			continue
		}
		queue = append(queue, n)
	}
	return queue
}

func alreadyInQueue(queue []*Node, node *Node) bool {
	for _, n := range queue {
		if n == node {
			return true
		}
	}
	return false
}

func inNeighbor(neighbors []*Node, searchNode *Node) bool {
	for _, n := range neighbors {
		if n == searchNode {
			return true
		}
	}
	return false
}

func main() {
	var n1 Node
	var n2 Node
	var n3 Node

	n1.FirstName = "arc"
	n2.FirstName = "bob"
	n3.FirstName = "unkani"

	var twitter SocialGraph

	twitter.AddNode(&n1)
	twitter.AddNode(&n2)
	twitter.AddNode(&n3)

	twitter.AddEdge(&n1, &n3) // arc and unkani
	twitter.AddEdge(&n2, &n3) // unkani and bob

	fmt.Println("Are arc and bob connected?", twitter.IsRelated(&n1, &n2))
}
