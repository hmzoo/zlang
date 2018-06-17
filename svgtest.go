package main

import (
  "github.com/ajstarks/svgo"
"os"
)

type Pos struct{
  canvas *svg.SVG
  X int
  Y int
  }

  func NewPos(f *os.File) *Pos {
    p := &Pos{canvas:svg.New(f),X:0,Y:0}
    return p
  }

  func (p *Pos) Start(w,h int) {
    p.canvas.Start(w,h)
  }
  func (p *Pos) End() {
    p.canvas.End()
  }

  func (p *Pos) Move(x,y int) {
    p.X=x
    p.Y=y
  }

  func (p *Pos) Node(t string) {
    p.canvas.Circle(p.X, p.Y, 30,"fill:white;stroke:black")
    p.canvas.Text(p.X, p.Y, t, "text-anchor:middle;font-size:16px;fill:black")
  }

  func (p *Pos) Nud(t string) {
    npx:=50
    npy:=50
    p.canvas.Line(p.X,p.Y,p.X-npx,p.Y+npy,"fill:none;stroke:black")
    p.canvas.Circle(p.X-npx, p.Y+npy, 30,"fill:white;stroke:black")
    p.canvas.Text(p.X-npx, p.Y+npy, t, "text-anchor:middle;font-size:16px;fill:black")
  }

  func (p *Pos) Led(t string) {
    npx:=50
    npy:=50
    p.canvas.Line(p.X,p.Y,p.X+npx,p.Y+npy,"fill:none;stroke:black")
    p.canvas.Circle(p.X+npx, p.Y+npy, 30,"fill:white;stroke:black")
    p.canvas.Text(p.X+npx, p.Y+npy, t, "text-anchor:middle;font-size:16px;fill:black")
  }

func main(){
  fo, err := os.Create("output.svg")
  if err != nil {
      panic(err)
  }
  p:=NewPos(fo)
  p.Start(1000,1000)
  p.Move(200,200)
  //p.Node("hello")
  p.Nud("nud")
  p.Led("led")
  p.Node("hello")
  p.End()


}
