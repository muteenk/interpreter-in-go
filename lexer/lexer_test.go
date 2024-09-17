package lexer

import (
  "testing"
  "monkey/token"
)


func TestNewToken(t *testing.T){
  input := `let a = 2;
let b = 3;

let add = fn(a, b) {
  a + b;
};

let result = add(a, b);`; // input string

  tests := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.LET, "let"},
    {token.IDENT, "a"},
    {token.ASSIGN, "="},
    {token.INT, "2"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENT, "b"},
    {token.ASSIGN, "="},
    {token.INT, "3"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENT, "add"},
    {token.ASSIGN, "="},
    {token.FUNCTION, "fn"},
    {token.LPAREN, "("},
    {token.IDENT, "a"},
    {token.COMMA, ","},
    {token.IDENT, "b"},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.IDENT, "a"},
    {token.PLUS, "+"},
    {token.IDENT, "b"},
    {token.SEMICOLON, ";"},
    {token.RBRACE, "}"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENT, "result"},
    {token.ASSIGN, "="},
    {token.IDENT, "add"},
    {token.LPAREN, "("},
    {token.IDENT, "a"},
    {token.COMMA, ","},
    {token.IDENT, "b"},
    {token.RPAREN, ")"},
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
