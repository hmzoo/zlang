package vtree

import ()

//USEFULL

func (n *Node) GetLevel() int {
	i := 0
	for f := n.GetParent(); f != nil; f = f.GetParent() {
		i = i + 1
	}
	return i
}

func (n *Node) GetDepth() int {
	r := 0
	l := 0
	if n.left != nil {
		l = 1 + n.left.GetDepth()
	}
	if n.right != nil {
		r = 1 + n.right.GetDepth()
	}
	if r > l {
		return r
	}
	return l
}

func (n *Node) CountLR() int {
	i := 0
	if n.HasLeft() {
		i = i + 1
	}
	if n.HasRight() {
		i = i + 1
	}
	return i
}

func (n *Node) GetRoot() *Node {

	if !n.IsRoot() {
		return n.GetParent().GetRoot()
	}
	return n

}

func (n *Node) CountLevel() int{
  r:=n.GetRoot()
  return r.countLevel(n.GetLevel()-1)
}

func (n *Node) countLevel(i int) int {
t:=0
if(i!=0){
  if(n.HasLeft()){t=t+n.left.countLevel(i-1)}
  if(n.HasRight()){t=t+n.right.countLevel(i-1)}
}else{
  return n.CountLR()
}
return t
}
