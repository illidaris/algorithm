package recursive

import (
	"strings"
)

func Recursive(root INode, nodes []INode) {
	for _, node := range nodes {
		if node.GetParentID() == root.GetID() {
			root.AddChild(node)
		}
	}
	if children := root.GetChildren(); len(children) > 0 {
		for _, child := range root.GetChildren() {
			Recursive(child, nodes)
		}
	}
}

func Flat(root INode, nodes *[]INode) {
	*nodes = append(*nodes, root)
	if children := root.GetChildren(); len(children) > 0 {
		for _, child := range root.GetChildren() {
			Flat(child, nodes)
		}
	}
}

func ParseLinesToNodes(sep string, f func() INameNode, lines ...string) map[int]map[string]INode {
	nodesArr := map[int]map[string]INode{}
	c := counter(0)
	for _, line := range lines {
		words := strings.Split(line, sep)
		for index, word := range words {
			if _, ok := nodesArr[index]; !ok {
				nodesArr[index] = map[string]INode{}
			}
			if _, ok := nodesArr[index][word]; ok {
				continue
			}
			n := f()
			n.SetID(c())
			n.SetName(word)
			n.SetRaw(line)
			nodesArr[index][word] = n
			if index == 0 {
				continue
			}
			if parent, ok := nodesArr[index-1][words[index-1]]; ok {
				n.SetParentID(parent.GetID())
			}
		}
	}
	return nodesArr
}

// counter 闭包计数器 （线程不安全）
func counter(id int64) func() int64 {
	return func() int64 {
		id += 1
		return id
	}
}
