package token


type TokenType string

type Token struct {
  Type TokenType
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
