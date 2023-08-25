package recursive

type INode interface {
	GetID() int64
	GetParentID() int64
	AddChild(node ...INode)
	GetChildren() []INode
}

type INameNode interface {
	INode
	SetID(id int64)
	SetParentID(pid int64)
	SetName(name string)
	SetRaw(raw string)
}

var _ = INameNode(&Node{})

type Node struct {
	Id       int64
	ParentId int64
	Name     string
	Raw      string
	Childs   []INode
}

func (n *Node) GetID() int64 {
	return n.Id
}
func (n *Node) GetParentID() int64 {
	return n.ParentId
}
func (n *Node) AddChild(node ...INode) {
	n.Childs = append(n.Childs, node...)
}
func (n *Node) GetChildren() []INode {
	return n.Childs
}

func (n *Node) SetID(id int64) {
	n.Id = id
}

func (n *Node) SetParentID(pid int64) {
	n.ParentId = pid
}
func (n *Node) SetName(name string) {
	n.Name = name
}
func (n *Node) SetRaw(raw string) {
	n.Raw = raw
}
