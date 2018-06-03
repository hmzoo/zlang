package main

import (
	"fmt"
)

var index int
var level int
var formule []string

func GetBP() int {
	s := GetLit()
	p := 0
	if s == "+" {
		p = 10
	}
	if s == "-" {
		p = 10
	}
	if s == "*" {
		p = 20
	}
	if s == "/" {
		p = 20
	}
	if s == "%" {
		p = 40
	}
	if s == "^" {
		p = 40
	}
	return p
}

func main() {
	f := "+A+B*(C+(D-E)/F)^G^H"

  fmt.Println(f)
	for x := 0; x < len(f); x = x + 1 {
		formule = append(formule, string([]rune(f)[x]))
	}
	//formule = []string{"+", "A", "+", "B", "*", "(", "-", "C", "+", "D", ")", "-", "E", "/", "F", "^", "G", "^", "H"}
	index = 0
	level = 0
  fmt.Println()
	fmt.Println(Expression(0))
}

func Next() {
	if index < len(formule)-1 {
		index = index + 1
	}
}

func IsEnd() bool {
	return index+1 < len(formule)
}

func GetLit() string {
	return formule[index]
}

func Nud() string {

	if GetLit() == "(" {
    l:=0
    s:=""
    Next()
		for GetLit() != ")" || l!=0 {
      if GetLit() == "(" {l=l+1}
      if GetLit() == ")" {l=l-1}
    s=s+GetLit()
    Next()
    }

    return "(" + s + ")"
	}


	if GetLit() == "-" {
		Next()
		return "'-" + GetLit() + "'"
	}
	if GetLit() == "+" {
		Next()
		return "'" + GetLit() + "'"
	}
	r := "'" + GetLit() + "'"
	return r
}

func Led(left string) string {
	o := GetBP()
	l := GetLit()
	Next()

	if l == "^" {
		return "[" + left + " " + l + " " + Expression(o-1) + "]"
	}
	return "[" + left + " " + l + " " + Expression(o) + "]"

}

func Expression(v int) string {
	left := Nud()

	Next()
	for IsEnd() && v+level < GetBP() {
		left = Led(left)
		Next()
	}

	return left

}
