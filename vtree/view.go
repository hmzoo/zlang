package vtree

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
  "strconv"
)

type View struct {
	svg  *svg.SVG
	tree *Tree
}

func NewView(t *Tree) *View {
	v := &View{tree: t}
	return v
}

func (v *View) Draw(f string) {
	fo, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	v.svg = svg.New(fo)
	v.svg.Start(1000, 1000)
  root := v.tree.GetRoot()
  v.drawLine(root, 60, 60)
	v.drawNode(root, 500, 300)
	v.svg.End()
}

func (v *View) drawLine(n *Node, x int, y int) {

  //fmt.Println(n.GetLevel(),n.GetDepth(),ratio,50*ratio)
  xd,yd :=v.getNXY(n)

	if n.GetRight() != nil {
    xdr,ydr :=v.getNXY(n.GetRight())
    v.svg.Line(xdr,ydr,xd,yd,"fill:none;stroke:red")
		v.drawLine(n.GetRight(), xd, yd)
	}
	if n.GetLeft() != nil {
    xdl,ydl :=v.getNXY(n.GetLeft())
    v.svg.Line(xdl,ydl,xd,yd,"fill:none;stroke:black")
		v.drawLine(n.GetLeft(), xd, yd)
	}
}

func (v *View) drawNode(n *Node, x int, y int) {
//  var ratio float64

  //fmt.Println(n.GetLevel(),n.GetDepth(),ratio,50*ratio)
  xd,yd :=v.getNXY(n)

	v.svg.Circle(xd, yd, 30, "fill:white;stroke:black")
	v.svg.Text(xd, yd, n.GetLabel(), "text-anchor:middle;font-size:20px;fill:black")
  v.svg.Text(xd-20, yd+12, strconv.Itoa(n.GetLX()), "text-anchor:middle;font-size:16px;fill:black")
  v.svg.Text(xd+20, yd+12, strconv.Itoa(n.GetLY()), "text-anchor:middle;font-size:16px;fill:black")
  v.svg.Text(xd, yd+16, strconv.Itoa(n.GetIndex()), "text-anchor:middle;font-size:16px;fill:black")
	if n.GetRight() != nil {
		v.drawNode(n.GetRight(), xd, yd)
	}
	if n.GetLeft() != nil {
		v.drawNode(n.GetLeft(), xd, yd)
	}
}

func (v *View) getNXY(n *Node) (int,int){
  w:=v.tree.GetWidthLevel(n.GetLY())
  fmt.Println(n.GetLY(),w)
  dx:=int(800/float64(w+1))
  xd:=dx+n.GetLX()*dx+60
  yd:=n.GetLY()*120+60
  return xd,yd
}
