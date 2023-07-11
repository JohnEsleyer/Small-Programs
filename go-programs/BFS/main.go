// Algorithm
// 1. Start at the root node
// 2. Add the root node to the queue
// 3. While the queue is not empty
//		1. Remove the first node from the queue
//		2. Mark the node visited
//		3. Add the node's neighbors to the queue
//	4. Repeat steps 3 and 4 until the queue is empty

package main

import (
	"fmt"
)

type Node struct {
	value    int
	children []*Node
	visited  bool
}

func BFS(root *Node) {
	queue := []*Node{}
	queue = append(queue, root)
	root.visited = true

	for len(queue) != 0 {
		// Dequeue a node from the queue
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", current.value)

		// Enqueue unvisited child nodes
		for _, child := range current.children {
			if !child.visited {
				queue = append(queue, child)
				child.visited = true
			}
		}
	}
}

func main() {
	// Create a sample graph nodes and edges
	node1 := &Node{value: 1, visited: false}
	node2 := &Node{value: 2, visited: false}
	node3 := &Node{value: 3, visited: false}
	node4 := &Node{value: 4, visited: false}
	node5 := &Node{value: 5, visited: false}

	node1.children = append(node1.children, node2)
	node1.children = append(node1.children, node3)
	node2.children = append(node2.children, node4)
	node3.children = append(node3.children, node4)
	node3.children = append(node3.children, node5)

	// Perform BFS traversal starting from node1
	fmt.Println("BFS Traversal:")
	BFS(node1)
}
