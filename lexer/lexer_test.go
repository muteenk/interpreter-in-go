package lexer

import (
  "testing"
  "monkey/token"
)


func TestNewToken(t *testing.T){
  input := "=+(){},;"; // input string

  tests := []struct {
    Type token.TokenType
    Literal string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.EOF, ""}  
  }


  
}
