package main

import(
  "github.com/zlang/vtree"
  //"fmt"
)

func main(){
/*
  t:=vtree.NewTree()
  t.NROOT("ROOT").NL("LEFT").NR("RIGHT").ML().NL("SUBL").NR("SUBR")

  t.HorFix()

  n := t.MROOT().GetCurrent()
  //v:=vtree.NewView(n)
  //v.Draw("tree.svg")
*/
  ht:=vtree.NewTree()
  ht.NFIRST("3",0).NN("*",20).NN("10",0).NN("+",10).NN("4",0).NN("*",20).NN("5",0)
  ht.NN("/",20).NN("8",0).NN("+",10).NN("6",0).NN("*",20).NN("3",0).NN("+",10).NN("9",0)
  ht.NN("/",20).NN("3",0).NN("*",20).NN("4",0)
  ht.VerFix()
//  ht.MarkPos()
  
  vv:=vtree.NewView(ht)
  vv.Draw("tree2.svg")
}
