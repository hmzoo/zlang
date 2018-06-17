package main

import(
  "fmt"
  "github.com/zlang/etree"
  "github.com/zlang/lex"
)


func main(){
  s:="-A+B*(C+(D-12)/F)^G^H"

  lex:=lex.NewLex(s)
  toks:=lex.ReadToks()


  e:=etree.NewETree(toks)

  fmt.Println(s)
  fmt.Println(toks)
  fmt.Println(e.Eval())


}
