package lex

import (
  //"fmt"
  "strconv"
)

type Lex struct {
	toks   []Tok
	input  []rune
	pos    int
	eof bool
	cur    string
	curd   string
	curw   string
}

func NewLex(text string) *Lex {
	l := &Lex{toks: []Tok{}, pos: 0, eof: false, cur: "", curd: "", curw: "", input: []rune(text)}
	return l
}

func (l *Lex) addTok(t Tok) *Lex {
	l.toks = append(l.toks, t)
	l.clw()
	return l
}

func (l *Lex) clw() *Lex {
	l.curw = ""
	return l
}

func (l *Lex) read() *Lex {


	if len(l.input) > l.pos+1 {
		l.cur = string(l.input[l.pos])
		l.curd = l.cur + string(l.input[l.pos+1])
	}

	if len(l.input) == l.pos+1 {
		l.cur = string(l.input[l.pos])
		l.curd = ""
	}

	if len(l.input) == l.pos {
		l.cur = ""
		l.curd = ""
	}

	l.pos = l.pos + 1

	return l
}

func (l *Lex) keep() *Lex {
  l.curw = l.curw + l.cur
  return l
}

func (l *Lex) checkEOF() *Lex {
	if l.cur == "" {
		l.addTok(Tok{t: EOF})
    l.eof=true
	}
	return l
}

func (l *Lex) checkSpace() *Lex {
	if IsBlank(l.cur) {
    l.checkWord()
    for l.read();IsBlank(l.cur);l.read(){}
	}
	return l
}

func (l *Lex) checkQuote() *Lex {
	if l.cur == "\"" {
    l.checkWord()
    for l.read().clw();l.cur!="\""&&l.cur!="";l.read(){l.keep()}
    l.addTok(Tok{t: STRING, lit: l.curw}).read()
	}
	return l
}

func (l *Lex) checkDouble() *Lex {
	if l.curd != "" {
		if value, exist := AllToks[l.curd]; exist {
			return l.checkWord().addTok(value).read()
		}
	}
	return l
}

func (l *Lex) checkSimple() *Lex {
	if l.cur != "" {
    //fmt.Print (l.cur+" ")
		if value, exist := AllToks[l.cur]; exist {
      //fmt.Println (value)
			return l.checkWord().addTok(value)
		}
    l.keep()
    //fmt.Println ("...")
	}
	return l
}

func (l *Lex) checkWord() *Lex {
  s:=l.curw
  //check  int
  _,err := strconv.Atoi(s)
  if(err==nil){
    return l.addTok(Tok{t:INT,lit:s})
  }
  //check  float
  _,err = strconv.ParseFloat(s,64)
  if(err==nil){
    return l.addTok(Tok{t:FLOAT,lit:s})
  }

  if value, exist := AllToks[s]; exist {
    return l.addTok(value)
  }

	if l.curw != "" && l.curw != " " {
		return l.addTok(Tok{t: IDENT, lit: l.curw})
	}
	return l
}



func (l *Lex) ReadToks() []Tok {
	for l.read(); !l.eof ; l.read() {
			l.checkEOF().checkSpace().checkQuote().checkDouble().checkSimple()
	}

	return l.toks

}
