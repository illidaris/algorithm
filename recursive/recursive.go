package recursive

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
