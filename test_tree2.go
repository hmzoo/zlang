package main

import(
  "github.com/zlang/vtree"
  //"fmt"
)

func main(){

  t:=vtree.NewTree()
  t.NROOT("COL").NL("LT").ML()
  t.NL("LTP").ML().NL("+")
  t.HorFix()

//  n := t.MROOT().GetCurrent()
  //v:=vtree.NewView(n)
  //v.Draw("tree.svg")

//  ht:=vtree.NewTree()
  //ht.NFIRST("COLLEGE",40).NN("LTP",20).NN("IN",10).NN("*",20).NN("N2",0).NN("+",10).NN("LTP",0)
//  ht.NN("+",10).NN("LTS",0).NN("+",10).NN("N2",0).NN("*",20).NN("N2",0).NN("*",20).NN("N2",0)
//  ht.NFIRST("N3",1).NN("WITH",2).NN("N2",3).NN("WITH",4).NN("N2",5).NN("IN",6).NN("LTP",0).NN("IN",20).NN("COL",40)
 //ht.NN("IN",40).NN("N2",0).NN("WITH",30).NN("N2",0)
//  ht.VerFix()
//  ht.MarkPos()

  vv:=vtree.NewView(t)
  vv.Draw("tree2.svg")
}
