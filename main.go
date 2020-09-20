package main

import "fmt"

type Node struct {
	path     string
	children []*Node
}

func main() {

	root := Node{
		path: "te",
		children: []*Node{
			&Node{
				path: "am",
			},
			&Node{
				path: "st",
			},
		},
	}

	node, exactMatch := findNode(&root, "abcd")

	if node != nil {
        fmt.Printf("Found a matching node: %v with exact match: %v", node, exactMatch)
	} else {
	    fmt.Printf("Not found")
    }
}

func (node *Node) addNode(path string) error {

    return nil
}

func findNode(node *Node, path string) (*Node, bool) {
	nodePath := node.path

	if !AreLettersEqual(nodePath, path, 0) {
		return nil, false
	}

	minLength := shortestString(nodePath, path)

	for i := range minLength {
		if AreLettersEqual(nodePath, path, i) {
			if i == (len(path) - 1) {
				return node, true
			}

			if i == len(nodePath)-1 {
				if len(node.children) == 0 {
					return node, false
				}

				for _, child := range node.children {
					if foundNode, exactMatch := findNode(child, path[i+1:]); foundNode != nil {
						return foundNode, exactMatch
					}
				}
			}
		} else {
			return node, false
		}
	}

	return nil, false
}



func AreLettersEqual(path string, path2 string, i int) bool {
    if i >= len(path) || i >= len(path2) {
        return false
    }

	letterA := string(path[i])
	letterB := string(path2[i])

	return letterA == letterB
}

func shortestString(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b
}
