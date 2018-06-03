package etree
import(
  "github.com/zlang/lex"
)

type ETree struct {
	formule []lex.Tok
	index   int
}

func NewETree(f []lex.Tok) *ETree {
	e := &ETree{index: 0}
  e.formule = f
	return e
}

func (e *ETree) Eval() string{
  e.index=0
  return e.expression(0)
}

func (e *ETree) next() {
	e.index = e.index + 1
}

func (e *ETree) isNotLast() bool {
	return e.index+1 < len(e.formule)
}

func (e *ETree) getLit() string {
	return e.formule[e.index].GetLit()
}

func (e *ETree) getType() string {
	return e.formule[e.index].GetType()
}

func (e *ETree) getBP() int {
	s := e.formule[e.index].GetType()
	p := 0
	if s == lex.PLUS || s == lex.MINUS {
		p = 10
	}
	if s == lex.MULTIPLY || s == lex.DIVIDE {
		p = 20
	}
	if s == lex.POWER {
		p = 30
	}
	if s == lex.MODULO {
		p = 40
	}
	return p
}

func (e *ETree) nud() string {

	if e.getType() == lex.LPARENT {
    l:=0
    s:=[]lex.Tok{}
    e.next()
		for e.getType() != lex.RPARENT || l!=0 {
      if e.getType() == lex.LPARENT {l=l+1}
      if e.getType() == lex.RPARENT {l=l-1}
    s=append(s,e.formule[e.index])
    e.next()
    }

    return "(" + NewETree(s).Eval() + ")"
	}


	if e.getType() == lex.MINUS {
		e.next()
		return "'-" + e.getLit() + "'"
	}
	if e.getType() == lex.PLUS {
		e.next()
		return "'" + e.getLit() + "'"
	}
	r := "'" + e.getLit() + "'"
	return r
}


func (e *ETree) led(left string) string {

  o := e.getBP()
	l := e.getType()
	e.next()

	if l == "^" {
		return "[" + left + " " + l + " " + e.expression(o-1) + "]"
	}
	return "[" + left + " " + l + " " + e.expression(o) + "]"

}

func (e *ETree) expression(v int) string {
	left := e.nud()

	e.next()
	for e.isNotLast() && v < e.getBP() {
		left = e.led(left)
		e.next()
	}

	return left

}
