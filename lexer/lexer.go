package lexer

import (
  "monkey/token"
)



type Lexer struct {

  input string
  position int
  ch byte

}



func New (input string) *Lexer{
  
  l := &Lexer{input: input}
  return l

}



