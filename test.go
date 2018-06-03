package main

import(
  "github.com/zlang/lex"
  "fmt"
)


func main(){


  
  text := "LET x =3+5*(10-3); \n PRINT \"oketo \";"
  l:=lex.NewLex(text)
  fmt.Print(text+"\n\n")

  fmt.Println(l.ReadToks())


}
