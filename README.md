# Practicing Interpreter Construction in GO

Writing an Interpreter just for fun

## Interpreter Tokens

```Go

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

```

## Interpreter Syntax, so far


```JavaScript

let a = 2;
let b = 32;

let add = fn(a, b) {
  a + b;
};

let result = add(a, b);

if (a < b) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;

```

