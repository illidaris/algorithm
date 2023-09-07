package recursive

import (
	"encoding/json"
	"strings"
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
	raws := []string{"中国/上海/上海/闵行", "中国/上海/上海/浦东", "中国/北京", "美国/加利福尼亚", "美国/华盛顿", "非洲", "英国/伦敦"}
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

func TestParseLinesToNodesWithSort(t *testing.T) {
	rawTxt := `学生/小学生
	学生/初中学生
	学生/高中学生
	学生/大学生
	学生/本硕在读学生
	机关、事业单位/公务员
	机关、事业单位/管理人员
	私营企业/员工
	私营企业/管理人员
	自我雇佣/商人
	自我雇佣/个体户
	自我雇佣/企业家
	自我雇佣/自媒体行业从业者
	艺术家/画家，雕塑家等传统艺术家
	艺术家/歌手，演员等
	艺术家/其他类别艺术家
	服务业从业者/餐饮、销售行业服务人员
	服务业从业者/快递、外卖行业服务人员
	服务业从业者/其他行业服务人员
	自由职业者
	其他
	无业`
	raws := strings.Split(rawTxt, "\n")
	//raws := []string{"中国/上海/上海/松江", "中国/上海/上海/青浦", "中国/北京", "美国/加利福尼亚", "美国/华盛顿", "非洲", "英国/伦敦"}
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
