package lex


const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	UNKNOWN     = "UNKNOWN"

	// IDENTIFIER + LITERALS
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	//operator
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
  MULTIPLY  = "*"
  DIVIDE  = "/"
  POWER  = "^"
	MODULO  = "%"

  //comparator
  BFALSE = "FALSE"
  BTRUE = "TRUE"
  AND ="&&"
  OR ="||"
  EQ ="=="
  NOTEQ ="!="
  LESS ="<"
  MORE =">"
  LESSEQ ="<="
  MOREEQ =">="


	//delimiters
  NL = "NL"
  QUOTE   = "\""
	PIPE   = "|"
	COMMA   = ","
  SEMICOLON   = ";"
  LPARENT = "("
	RPARENT = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LHOOK  = "["
  RHOOK  = "]"

  //keywords
  LET ="LET"
  IF ="IF"

  //func
  PRINT ="PRINT"
)
