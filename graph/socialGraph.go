package main

import "fmt"

type SocialGraph struct {
	nodes []*Node
	edges map[Node][]*Node
}

type Node struct {
	person Person
}

type Person struct {
	FirstName string
	LastName  string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.person)
}

func (g *SocialGraph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *SocialGraph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *SocialGraph) IsRelated(mainNode, otherNode *Node) bool {
	for _, node := range g.edges[*mainNode] {
		if node == otherNode {
			return true
		}
	}
	return false
}

func main() {
	p1 := Person{"arc", "thefallen"}
	p2 := Person{"bob", "620"}

	n1 := Node{p1}
	n2 := Node{p2}

	var twitter SocialGraph
	twitter.AddNode(&n1)
	twitter.AddNode(&n2)
	twitter.AddEdge(&n1, &n2)

	fmt.Println("Are bob and arc friends?", twitter.IsRelated(&n1, &n2))
}
