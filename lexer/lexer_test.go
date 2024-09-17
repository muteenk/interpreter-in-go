package lexer

import (
  "testing"
  "monkey/token"
)


func TestNewToken(t *testing.T){
  input := "=+(){},;"; // input string

  tests := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.EOF, ""}, 
  }

  l := newLexer(input)

  for i, tt := range tests {
    tok := l.nextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
    }
  }
  
}
