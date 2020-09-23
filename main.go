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
				children:  []*Node {
					&Node {
						path: "ify",
					},
				},
			},
		},
	}

	node, indexInNode, indexInPath := findNode(&root, "abcd", 0)

	if node != nil {
        fmt.Printf("Found a matching node: %v with index in node: %d, index in path: %d", node, indexInNode, indexInPath)
	} else {
	    fmt.Printf("Not found")
    }
}

func (node *Node) addNode(path string) error {

    return nil
}

func findNode(node *Node, path string, currentIndex int) (*Node, int, int) {
	nodePath := node.path

	currentPath := path[currentIndex:]

	if !AreLettersEqual(nodePath, currentPath, 0) {
		return nil, -1, -1
	}

	minLength := shortestString(nodePath, currentPath)

	for i := range minLength {
		if AreLettersEqual(nodePath, currentPath, i) {
			if i == (len(currentPath) - 1) {
				return node, i, currentIndex + i
			}

			if i == len(nodePath)-1 {
				if len(node.children) == 0 {
					return node, i, currentIndex + i
				}

				for _, child := range node.children {
					if foundNode, indexInNodePath, indexInOriginalPath := findNode(child, path, i+1+currentIndex); foundNode != nil {
						return foundNode, indexInNodePath, indexInOriginalPath
					}
				}
			}
		} else {
			return node, i, currentIndex
		}
	}

	return nil, -1, -1
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
