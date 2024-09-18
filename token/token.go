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
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
  
  // Identifiers & literals
  IDENT = "IDENTIFIER"
  INT = "INTEGER"
  
  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  ASTERISK = "*"
  SLASH = "/"
  NOT = "!"
  
  LT = "<"
  GT = ">"
  
  EQ = "=="
  NOT_EQ = "!="

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
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}


func LookupIdentifier(ident string) TokenType {
  tok, ok := keywords[ident]

  if ok {
    return tok
  }

  return IDENT
}



