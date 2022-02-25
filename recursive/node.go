package recursive

type INode interface {
	GetID() int64
	GetParentID() int64
	AddChild(node ...INode)
	GetChildren() []INode
}
