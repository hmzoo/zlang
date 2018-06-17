package main

import(
  "github.com/zlang/vtree"
  "fmt"
)

func main(){

  n1 := vtree.NewNode("n1")
  n2 := vtree.NewNode("n2")
  n3 := vtree.NewNode("n3")
  n4 := vtree.NewNode("n4")
  n5 := vtree.NewNode("n5")
  n6 := vtree.NewNode("n6")
  n7 := vtree.NewNode("n7")
  n8 := vtree.NewNode("n8")
  n9 := vtree.NewNode("n9")
  n1.SetLeft(n2)
  n1.SetRight(n3)
  n3.SetRight(n4)
  n4.SetLeft(n5)
  n4.SetRight(n6)
  n2.SetLeft(n7)
  n2.SetRight(n8)
  n6.SetRight(n9)

  fmt.Println(n2.GetBP())

  fmt.Println(n1.GetWidth(),n1.GetHeight())

  v:=vtree.NewView(n1)
  v.Draw("tree.svg")
  n1.RangeInFixe()
}
