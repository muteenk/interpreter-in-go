package token

type Token struct {
  Type string
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  LET = "LET"
  FUNCTION = "FUNCTION"
  
  IDENT = "IDENTIFIER"
  INT = "INTEGER"
  
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  COMMA = ","
  SEMICOLON = ";"
);
