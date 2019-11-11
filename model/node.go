package model

type NodeMask uint32
type NodePath []*Node

const (
	NODE_T string = "osg::Node"
)

type ComputeBoundingSphereCallback struct {
	Object
}

type Node struct {
	Object
	CullingActive bool
	NodeMask      NodeMask
	Dscriptions   []string
	InitialBound  Sphere3f
	States        *StateSet
	Parents       []*Group

	Callback       *ComputeBoundingSphereCallback
	UpdateCallback *Callback
	EventCallback  *Callback
	CullCallback   *Callback
}

func NewNode() Node {
	obj := NewObject()
	obj.Type = NODE_T
	return Node{Object: obj}
}

func (n *Node) Accept(nv *NodeVisitor) {
	if nv.ValidNodeMask(n) {
		nv.PushOntoNodePath(n)
		nv.Apply(n)
		nv.PopFromNodePath(n)
	}
}
func (n *Node) Ascend(nv *NodeVisitor) {

}

func (n *Node) Traverse(nv *NodeVisitor) {

}
