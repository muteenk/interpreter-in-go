package lexer

import (
  "monkey/token"
)


type Lexer struct {

  input string
  inputSize int
  readCursor int
  cursor int
  ch byte

}


// Lexer Constructor
func newLexer (input string) *Lexer{
  lex := &Lexer{input: input, inputSize: len(input)}
  lex.readChar()
  return lex
}


// Advance to the next character
func (lex *Lexer) readChar() {
  
  if lex.readCursor >= lex.inputSize {
    lex.ch = 0;
  } else {
    lex.ch = lex.input[lex.readCursor]
  }

  lex.cursor = lex.readCursor
  lex.readCursor++; 

}


// Generates Next Token based on characters
func (lex *Lexer) nextToken() token.Token {
  var tok token.Token
  
  switch lex.ch {
    
    case '=':
      tok = newToken(token.ASSIGN, lex.ch)
    case ';':
      tok = newToken(token.SEMICOLON, lex.ch)
    case '(':
      tok = newToken(token.LPAREN, lex.ch)
    case ')':
      tok = newToken(token.RPAREN, lex.ch)
    case '{':
      tok = newToken(token.LBRACE, lex.ch)
    case '}':
      tok = newToken(token.RBRACE, lex.ch)
    case ',':
      tok = newToken(token.COMMA, lex.ch)
    case '+':
      tok = newToken(token.PLUS, lex.ch)
    case 0:
      tok.Literal = ""
      tok.Type = token.EOF
  }

  lex.readChar()
  return tok
}


// Token Creator
func newToken(tokenType token.TokenType, ch byte) token.Token{
  return token.Token{Type: tokenType, Literal: string(ch)}

}


