package main

import (
	"fmt"
  "strconv"
)

var LBPS map[string]int

func init(){
  LBPS =map[string]int{"+":10,"-":20}

}

type Tok struct {
  lit string
  lbp int
  next *Tok
  op Op
}

func (t *Tok) Next() *Tok {
  return t.next
}

func (t *Tok) GetLit() string {
  return t.lit
}

func (t *Tok) GetLBP() int {
  return t.lbp
}

func (t *Tok) SetOp(op Op)  {
  op.SetTok(t)
  t.op=op
}

func (t *Tok) NextTok(s string) *Tok{
  p:=0
  if val,ok := LBPS[s]; ok{
    p=val
  }
  tok := &Tok{lit:s,lbp:p}
  t.next = tok
  return tok
}


type Op interface {
  SetTok(*Tok)
  Nud() float64
  Led() float64
}

//LIT
type Op_lit struct {
  Tok *Tok
}

func (op Op_lit) SetTok(tok *Tok) {
  op.Tok =tok
}
func (op Op_lit) Nud() float64{
  f,err := strconv.ParseFloat(op.Tok.GetLit(), 64)
  if(err!=nil){return f}
  return 0
}
func (op Op_lit) Led() float64{
  f,err := strconv.ParseFloat(op.Tok.GetLit(), 64)
  if(err!=nil){return f}
  return 0
}

//Add
type Op_Add struct {
  Tok *Tok
  lbp int
}

func (op Op_Add) SetTok(tok *Tok) {
  op.Tok =tok
  lbp=10
}
func (op Op_Add) Nud() float64{
  f,err := strconv.ParseFloat(op.Tok.GetLit(), 64)
  if(err!=nil){return f}
  return 0
}
func (op Op_Add) Led() float64{
  return op.Tok.
  f,err := strconv.ParseFloat(op.Tok.GetLit(), 64)
  if(err!=nil){return f}
  return 0
}

func main() {
	tok := &Tok{lit:"1"}
  tok.SetOp(Op_lit{})
  tok.NextTok("+").NextTok("3")

  for t := tok; t!=nil;t=t.Next(){
    fmt.Println(t.GetLit(),t.GetLBP())
  }

}





/*
func Expression(rbp int){
  token := Next()
  left := t.nud()
  for r :=rbp ;r < token.lbp; {
    t := token
    token= next()
    left = t.led(left)
  }
  return left

}
*/
