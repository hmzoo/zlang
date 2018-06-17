package vtree

import (
	"fmt"
)

type Tree struct {
	current *Node
	cpt     int
	lcpt    []int
}

func NewTree() *Tree {
	return &Tree{cpt: 0}
}

func (t *Tree) GetCurrent() *Node {
	return t.current
}
func (t *Tree) SetCurrent(n *Node) {
	t.current = n
}

func (t *Tree) GetRoot() *Node {
	return t.MROOT().current
}

//EDITING
func (t *Tree) UPDTLABEL(s string) *Tree {
	if t.current == nil {
		return t
	}
	t.current.SetLabel(s)
	return t
}

func (t *Tree) UPDTBP(i int) *Tree {
	if t.current == nil {
		return t
	}
	t.current.SetBP(i)
	return t
}

//TREE SHAPE

//VERTICAL BUILDING
func (t *Tree) NROOT(s string) *Tree {
	n := NewNode(s)
	t.SetCurrent(n)
	return t
}

func (t *Tree) NL(s string) *Tree {
	if t.current == nil {
		return t
	}

	n := NewNode(s)
	t.current.SetLeft(n)

	return t
}

func (t *Tree) NR(s string) *Tree {
	if t.current == nil {
		return t
	}

	n := NewNode(s)
	t.current.SetRight(n)

	return t
}

//VERTICAL MOVES

func (t *Tree) MROOT() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsRoot() {
		t.current = t.current.GetParent()
		return t.MROOT()
	}
	return t
}

func (t *Tree) MP() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsRoot() {
		t.current = t.current.GetParent()
	}
	return t
}

func (t *Tree) ML() *Tree {
	if t.current == nil {
		return t
	}
	if t.current.HasLeft() {
		t.current = t.current.GetLeft()
	}
	return t
}

func (t *Tree) MR() *Tree {
	if t.current == nil {
		return t
	}
	if t.current.HasRight() {
		t.current = t.current.GetRight()
	}
	return t
}

// VERTICAL REORDER

func (t *Tree) VerFix() {
	t.GFIRST()
	t.expr(0)
	t.MarkPos()
}

func (t *Tree) nud() *Node {
	n := t.current
	t.GN()
	return n
}

func (t *Tree) led(l *Node) *Node {
	n := t.current
	t.GN()
	n.SetLeft(l)
	n.SetRight(t.expr(n.GetBP()))
	return n
}

func (t *Tree) expr(v int) *Node {

	left := t.nud()

	for !t.current.IsLast() && v < t.current.GetBP() {
		left = t.led(left)
	}

	return left
}

//HORIZONTAL BUILDING
func (t *Tree) NFIRST(s string, bp int) *Tree {
	n := NewNode(s)
	n.SetBP(bp)
	n.SetIndex(0)
	t.SetCurrent(n)
	return t
}

func (t *Tree) NN(s string, bp int) *Tree {
	if t.current == nil {
		return t
	}

	n := NewNode(s)
	n.SetBP(bp)
	n.SetIndex(t.current.GetIndex() + 1)
	t.current.SetNext(n)
	t.current = n

	return t
}

//HORIZONTAL MOVES

func (t *Tree) GFIRST() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsFirst() {
		t.current = t.current.GetPrevious()
		return t.GFIRST()
	}

	return t
}

func (t *Tree) GLAST() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsLast() {
		t.current = t.current.GetNext()
		return t.GLAST()
	}

	return t
}

func (t *Tree) GN() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsLast() {
		t.current = t.current.GetNext()
	}
	return t
}

func (t *Tree) GP() *Tree {
	if t.current == nil {
		return t
	}
	if !t.current.IsFirst() {
		t.current = t.current.GetPrevious()
	}
	return t
}

// HORIZONTAL REORDER
func (t *Tree) HorFix() {
	r := t.MROOT().GetCurrent()
	t.SetCurrent(nil)
	t.infixe(r)
	i := 0
	for t.GFIRST(); !t.current.IsLast(); t.GN() {
		t.current.SetIndex(i)
		i = i + 1
	}
	t.current.SetIndex(i)
  t.MarkPos()
}

func (t *Tree) infixe(n *Node) {

	if n.left != nil {
		t.infixe(n.left)
	}
	if t.current != nil {

		t.current.SetNext(n)

		fmt.Println(n.GetLabel(), t.current.GetLabel(), t.current.GetIndex())
	} else {
		fmt.Println(n.GetLabel(), "NIL")
	}

	t.current = n

	if n.right != nil {
		t.infixe(n.right)
	}

}

//POSITION
func (t *Tree) MarkPos() {
	for i, _ := range t.lcpt {
		t.lcpt[i] = 0
	}
	t.MROOT().walky(0)
  t.lcpt[0]=1
	t.MROOT().walkx(0)

}
func (t *Tree) walky(y int) {
	n := t.current
	y = y + 1
	t.lcpt = append(t.lcpt, 0)
	if t.current.HasLeft() {
		t.ML().current.SetLY(y)
		t.walky(y)
	}
	t.current = n
	if t.current.HasRight() {
		t.MR().current.SetLY(y)
		t.walky(y)
	}
}

func (t *Tree) walkx(x int) {

	n := t.current

	t.cpt = t.cpt + 1
	x = x + 1

	if t.current.HasLeft() {

		t.ML().current.SetLX(t.lcpt[x])
		t.lcpt[x] = t.lcpt[x] + 1
		t.walkx(x)
	}
	t.current = n

	if t.current.HasRight() {

		t.MR().current.SetLX(t.lcpt[x])
		t.lcpt[x] = t.lcpt[x] + 1
		t.walkx(x)
	}
	t.current = n

}

func (t *Tree) GetWidthLevel(l int) int {
	if l < len(t.lcpt) {
		return t.lcpt[l]
	}
	return 0

}
