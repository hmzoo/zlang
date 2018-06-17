package vtree

import ()

type Node struct {
	label    string
	bp       int
  index       int
  lY       int
  lX       int
	left     *Node
	right    *Node
	parent   *Node
	next     *Node
	previous *Node
}

func NewNode(l string) *Node {
	return &Node{label: l, bp: 0}
}

func (n *Node) SetLabel(s string) {
	n.label = s
}

func (n *Node) GetLabel() string {
	return n.label
}

func (n *Node) SetBP(bp int) {
	n.bp = bp
}

func (n *Node) GetBP() int {
	return n.bp
}

func (n *Node) SetIndex(i int) {
	n.index = i
}

func (n *Node) GetIndex() int {
	return n.index
}

func (n *Node) SetLX(i int) {
	n.lX = i
}

func (n *Node) GetLX() int {
	return n.lX
}

func (n *Node) SetLY(i int) {
	n.lY = i
}

func (n *Node) GetLY() int {
	return n.lY
}

func (n *Node) SetLeft(node *Node) {
	n.left = node
  if(node.GetParent()!=n) {node.SetParent(n)}
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) SetRight(node *Node) {
	n.right = node
  if(node.GetParent()!=n) {node.SetParent(n)}
}

func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) SetParent(node *Node) {
	n.parent = node
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) SetNext(node *Node) {
  n.next = node
	if(node.GetPrevious()!=n) {node.SetPrevious(n)}

}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetPrevious(node *Node) {
	n.previous = node
  if(node.GetNext()!=n) {node.SetNext(n)}
}

func (n *Node) GetPrevious() *Node {
	return n.previous
}


func (n *Node) IsRight() bool {
	return n.parent != nil && n.parent.GetRight() == n
}

func (n *Node) IsLeft() bool {
	return n.parent != nil && n.parent.GetLeft() == n
}

func (n *Node) HasRight() bool {
	return n.right != nil
}
func (n *Node) HasLeft() bool {
	return n.left != nil
}



func (n *Node) IsRoot() bool {
	return n.parent == nil
}

func (n *Node) IsFirst() bool {
	return n.previous == nil
}

func (n *Node) IsLast() bool {
	return n.next == nil
}
