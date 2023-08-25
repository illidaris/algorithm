package recursive

import (
	"encoding/json"
	"testing"
)

type testNode struct {
	ID       int64
	ParentID int64
	Children []INode
}

func (n *testNode) GetID() int64 {
	return n.ID
}

func (n *testNode) GetParentID() int64 {
	return n.ParentID
}

func (n *testNode) AddChild(node ...INode) {
	if n.Children == nil {
		n.Children = make([]INode, 0)
	}
	n.Children = append(n.Children, node...)
}

func (n *testNode) GetChildren() []INode {
	return n.Children
}

var testNodes = []INode{
	&testNode{
		ID:       1,
		ParentID: 7,
	},
	&testNode{
		ID:       3,
		ParentID: 4,
	},
	&testNode{
		ID:       4,
		ParentID: 7,
	},
	&testNode{
		ID:       5,
		ParentID: 0,
	},
	&testNode{
		ID:       6,
		ParentID: 0,
	},
	&testNode{
		ID:       7,
		ParentID: 5,
	},
	&testNode{
		ID:       8,
		ParentID: 6,
	},
}

func TestRecursive(t *testing.T) {
	root := &testNode{
		ID:       0,
		ParentID: 0,
	}
	Recursive(root, testNodes)
	bs, err := json.Marshal(root)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(bs)
	println(jsonStr)
	con1 := root.Children[0].GetID() == 5
	con2 := root.Children[0].GetChildren()[0].GetID() == 7
	con3 := root.Children[1].GetID() == 6
	if !con1 || !con2 || !con3 {
		t.Error("no correct")
	}
	nodes := make([]INode, 0)
	Flat(root, &nodes)
	if len(nodes) != 8 {
		t.Error("len is not 8")
	}
}

func TestParseLinesToNodes(t *testing.T) {
	raws := []string{"中国/上海/上海/松江", "中国/上海/上海/青浦", "中国/北京", "美国/加利福尼亚", "美国/华盛顿", "非洲", "英国/伦敦"}
	nodesMap := ParseLinesToNodes("/", func() INameNode {
		return &Node{}
	}, raws...)
	ns := []INode{}
	for _, nodes := range nodesMap {
		for _, node := range nodes {
			ns = append(ns, node)
		}
	}
	root := &Node{}
	Recursive(root, ns)
	bs, _ := json.Marshal(root)
	println(string(bs))
}
