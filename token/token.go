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


var keywords = map[string] TokenType{
  "fn": FUNCTION,
  "let": LET,
}


func LookupIdentifier(ident string) TokenType {
  tok, ok := keywords[ident]

  if ok {
    return tok
  }

  return IDENT
}



