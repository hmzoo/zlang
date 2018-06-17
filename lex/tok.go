package lex

type Tok struct {
	t   string
	lit string
}

func (t *Tok) GetLit() string {
	return t.lit
}

func (t *Tok) GetType() string {
	return t.t
}

func (t *Tok) IsEOF() bool {
	return t.t == EOF
}

func (t *Tok) GetBP() int {
	s := t.GetType()
	p := 0
	if s == PLUS || s == MINUS {
		p = 10
	}
	if s == MULTIPLY || s == DIVIDE {
		p = 20
	}
	if s == POWER {
		p = 30
	}
	if s == MODULO {
		p = 40
	}

	return p
}

var AllToks map[string]Tok

func init() {

	AllToks = map[string]Tok{
		"\n":    Tok{t: NL},
		"!=":    Tok{t: NOTEQ},
		"==":    Tok{t: EQ},
		">=":    Tok{t: MOREEQ},
		"<=":    Tok{t: LESSEQ},
		">":     Tok{t: MORE},
		"<":     Tok{t: LESS},
		"=":     Tok{t: ASSIGN},
		"+":     Tok{t: PLUS},
		"-":     Tok{t: MINUS},
		"*":     Tok{t: MULTIPLY},
		"/":     Tok{t: DIVIDE},
		"(":     Tok{t: LPARENT},
		")":     Tok{t: RPARENT},
		"{":     Tok{t: LBRACE},
		"}":     Tok{t: RBRACE},
		"[":     Tok{t: LHOOK},
		"]":     Tok{t: RHOOK},
		"^":     Tok{t: POWER},
		",":     Tok{t: COMMA},
		";":     Tok{t: SEMICOLON},
		"LET":   Tok{t: LET},
		"IF":    Tok{t: IF},
		"PRINT": Tok{t: PRINT},
	}

}
