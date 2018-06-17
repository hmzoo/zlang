package lex

import (
	"fmt"
	"strconv"
  "math"
)

type Exp struct {
	toks []Tok
  pos int
  eof   bool
}

func NewExp(toks []Tok) *Exp {
	e := &Exp{pos: 0,eof:false, toks :toks}
	return e
}




func (e *Exp) next() *Tok{
  t := &e.toks[e.pos]
  e.pos=e.pos+1
  if e.pos < len(e.toks) {
		e.eof = false
	} else {
		e.eof = true
	}
  return t
}

func(e *Exp) nud() float64{
  t:=e.next()
  var result float64
  if(t.GetType()==INT || t.GetType()==FLOAT){
    result,_ =strconv.ParseFloat(t.GetLit(), 64)
  }
  return result
}

func(e *Exp) led(left float64) float64{
  var result float64
  t:=e.next()
  bp:=t.GetBP()

  if(t.GetType()==PLUS ){
    o:=e.Expr(bp)
    result=left+o
    fmt.Println(left,"+",o,result,bp)
  }

  if(t.GetType()==MINUS ){

    o:=e.Expr(bp)
    result=left-o
    fmt.Println(left,"-",o,result,bp)
  }

  if(t.GetType()==MULTIPLY ){
    o:=e.Expr(bp)
    result=left*o
    fmt.Println(left,"*",o,result,bp)
  }

  if(t.GetType()==DIVIDE ){
    o:=e.Expr(bp)
    result=left/o
    fmt.Println(left,"/",o,result,bp)
  }

  if(t.GetType()==POWER){
    o:=e.Expr(bp-1)
    result=math.Pow(left,o)
    fmt.Println(left,"^",o,result,bp)
  }

  if(t.GetType()==EOF ){
    result=left
    fmt.Println(left,"eof",result)
  }
  return result
}

func (e *Exp) Expr(v int) float64{

  left := e.nud()

  for !e.eof && v < e.toks[e.pos].GetBP() {
    left=e.led(left)
  }

  return left
}
