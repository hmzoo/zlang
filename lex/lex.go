package lex

import (
	"fmt"
	"strconv"
)

type Lex struct {
	input []rune
	pos   int
	eof   bool
}

func NewLex(text string) *Lex {
	l := &Lex{pos: 0, eof: false, input: []rune(text)}
	fmt.Println(len(text), len(l.input))
	return l
}

func (l *Lex) GetCur() string {
	if l.pos < len(l.input) {
		return string(l.input[l.pos])
	}
	return ""
}

func (l *Lex) GetPos() int {
	return l.pos
}

func (l *Lex) GetLen() int {
	return len(l.input)
}

func (l *Lex) GetCurN(n int) string {
	p := l.pos
	r := ""
	for i := p; i < n+p; i = i + 1 {
		if i < len(l.input) {
			r = r + string(l.input[i])
		}
	}
	return r
}

func (l *Lex) Eof() bool {
	return l.eof
}

func (l *Lex) Next() {
	l.pos = l.pos + 1
	if l.pos < len(l.input) {
		l.eof = false
	} else {
		l.eof = true
	}

}

func (l *Lex) Read() string {
	r := l.GetCur()
	l.Next()
	return r
}

//BLANK
func (l *Lex) NextIsBlank() bool {
	Blanks := []string{" ", "\t", "\r", "\n"}
	for _, v := range Blanks {
		if v == l.GetCur() {
			return true
		}
	}
	return false
}

func (l *Lex) EatBlank() {
	if l.Eof() {
		l.Next()
	}
	if l.NextIsBlank() {
		l.Next()
		l.EatBlank()
	}
}

//QUOTE
func (l *Lex) NextIsQuote() bool {
	if l.GetCur() == "\"" {
		return true
	}
	return false
}

func (l *Lex) ReadString(result string) string {
	r := l.Read()
	if r != "\"" {
		return l.ReadString(result + r)
	}
	return result
}

func (l *Lex) GetString() Tok {
	l.Next()
	return Tok{t: STRING, lit: l.ReadString("")}
}

//DOUBLE

func (l *Lex) NextIsDouble() bool {
	if _, exist := AllToks[l.GetCurN(2)]; exist {
		return true
	}
	return false
}

func (l *Lex) GetDouble() Tok {
	s := l.Read() + l.Read()
	if v, exist := AllToks[s]; exist {
		return v
	}
	return Tok{t: UNKNOWN, lit: s}
}

//DOUBLE

func (l *Lex) NextIsSimple() bool {
	if _, exist := AllToks[l.GetCur()]; exist {
		return true
	}
	return false
}

func (l *Lex) GetSimple() Tok {
	s := l.Read()
	if v, exist := AllToks[s]; exist {
		return v
	}
	return Tok{t: UNKNOWN, lit: s}
}

//LITERALS

func (l *Lex) IsLimit() bool {
	return l.Eof() || l.NextIsDouble() || l.NextIsSimple() || l.NextIsQuote() || l.NextIsBlank()
}

func (l *Lex) ReadLiteral(result string) string {
	r := l.Read()
	if !l.IsLimit() {
		return l.ReadLiteral(result + r)
	} else {
		return result + r
	}

}
func (l *Lex) GetLiteral() Tok {
	s := l.ReadLiteral("")

	_, err := strconv.Atoi(s)
	if err == nil {
		return Tok{t: INT, lit: s}
	}
	//check  float
	_, err = strconv.ParseFloat(s, 64)
	if err == nil {
		return Tok{t: FLOAT, lit: s}
	}

	if value, exist := AllToks[s]; exist {
		return value
	}

	return Tok{t: "IDENT", lit: s}

}

func (l *Lex) NextTok() Tok {
	if l.Eof() {
		return Tok{t: EOF}
	}
	if l.NextIsBlank() {
		l.EatBlank()
		return l.NextTok()
	}
	if l.NextIsQuote() {
		return l.GetString()
	}
	if l.NextIsDouble() {
		return l.GetDouble()
	}
	if l.NextIsSimple() {
		return l.GetSimple()
	}

	return l.GetLiteral()

}
