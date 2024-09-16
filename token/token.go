package token

type Token struct {
  Type string
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Keywords
  LET = "LET"
  FUNCTION = "FUNCTION"
  
  // Identifiers & literals
  IDENT = "IDENTIFIER"
  INT = "INTEGER"
  
  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"

  // Delimiters
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  COMMA = ","
  SEMICOLON = ";"
);
