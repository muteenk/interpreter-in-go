package main

import (
  "fmt"
  "monkey/lexer"
);

func main() {
  tok := token.Token{Type: token.LET, Literal: "let"};
  fmt.Println(tok); // {LET let}
}
