package lex

type Tok struct {
	t string
  lit string

}


func (t*Tok) GetLit() string {
  return t.lit
}

func (t*Tok) GetType() string {
  return t.t
}

func (t*Tok) IsEOF() bool {
  return t.t == EOF
}

var AllToks map[string]Tok
var Blanks []string

func IsBlank(s string) bool{
  for _,v := range Blanks {
    if(v==s){return true}
  }
  return false
}

func init(){
  Blanks =[]string{" ","\t","\r","\n"}
  AllToks =map[string]Tok{
    "\n" :Tok{t:NL},
    "!=" :Tok{t:NOTEQ},
    "==" :Tok{t:EQ},
    ">=" :Tok{t:MOREEQ},
    "<=" :Tok{t:LESSEQ},
    ">" :Tok{t:MORE},
    "<" :Tok{t:LESS},
    "=" :Tok{t:ASSIGN},
    "+" :Tok{t:PLUS},
    "-" :Tok{t:MINUS},
    "*" :Tok{t:MULTIPLY},
    "/" :Tok{t:DIVIDE},
    "(" :Tok{t:LPARENT},
    ")" :Tok{t:RPARENT},
    "^" :Tok{t:POWER},
    "," :Tok{t:COMMA},
    ";" :Tok{t:SEMICOLON},
    "LET" :Tok{t:LET},
    "IF" :Tok{t:IF},
    "PRINT" :Tok{t:PRINT},
  }


}
