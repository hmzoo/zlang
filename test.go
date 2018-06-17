package main

import (
	"fmt"
	"github.com/zlang/lex"

)

func main() {

	//text := "LET x =3+ 5*(10-3); { z!=3^4 ; \n {}PRINT \"oketo \"; 10.3 ; "
	text := "3+4^2^3-5*2/2+30"
	l := lex.NewLex(text)
	fmt.Print(text + "\n\n")

  toks:=[]lex.Tok{}

	for t := l.NextTok(); !t.IsEOF(); t = l.NextTok() {
		toks = append(toks,t)
	}
  toks = append(toks,l.NextTok())

  e := lex.NewExp(toks)

  fmt.Println(e)
  fmt.Println(e.Expr(0))



}
