package lexer

import (
  "testing"
  "monkey/token"
)


func TestNewToken(t *testing.T){
  input := `let a = 2;
let b = 32;

let add = fn(a, b) {
  a + b;
};

let result = add(a, b);
!-/*5;
5 < 10 > 5;

if (a < b) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;`; // input string

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
    {token.INT, "32"},
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
    {token.NOT, "!"},
    {token.MINUS, "-"},
    {token.SLASH, "/"},
    {token.ASTERISK, "*"},
    {token.INT, "5"},
    {token.SEMICOLON, ";"},
    {token.INT, "5"},
    {token.LT, "<"},
    {token.INT, "10"},
    {token.GT, ">"},
    {token.INT, "5"},
    {token.SEMICOLON, ";"},
    {token.IF, "if"},
    {token.LPAREN, "("},
    {token.IDENT, "a"},
    {token.LT, "<"},
    {token.IDENT, "b"},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RETURN, "return"},
    {token.TRUE, "true"},
    {token.SEMICOLON, ";"},
    {token.RBRACE, "}"},
    {token.ELSE, "else"},
    {token.LBRACE, "{"},
    {token.RETURN, "return"},
    {token.FALSE, "false"},
    {token.SEMICOLON, ";"},
    {token.RBRACE, "}"},
    {token.INT, "10"},
    {token.EQ, "=="},
    {token.INT, "10"},
    {token.SEMICOLON, ";"},
    {token.INT, "10"},
    {token.NOT_EQ, "!="},
    {token.INT, "9"},
    {token.SEMICOLON, ";"},
    {token.EOF, ""},
  }

  l := NewLexer(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
    }
  }
  
}
