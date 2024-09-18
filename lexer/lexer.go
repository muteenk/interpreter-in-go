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
func NewLexer (input string) *Lexer{
  lex := &Lexer{input: input, inputSize: len(input)}
  lex.readChar()
  return lex
}


// Generates Next Token based on characters
func (lex *Lexer) NextToken() token.Token {
  var tok token.Token

  lex.skipWhiteSpace()
  
  switch lex.ch {
    
    case '=':
      if lex.peekChar() == '=' {
        lex.readChar()
        tok = token.Token{Type: token.EQ, Literal: "=="} 
      } else {
        tok = newToken(token.ASSIGN, lex.ch)
      }
    case '!':
      if lex.peekChar() == '=' {
        lex.readChar()
        tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
      } else {
        tok = newToken(token.NOT, lex.ch)
      }
    case '+':
      tok = newToken(token.PLUS, lex.ch)
    case '-':
      tok = newToken(token.MINUS, lex.ch)
    case '/':
      tok = newToken(token.SLASH, lex.ch)
    case '*':
      tok = newToken(token.ASTERISK, lex.ch)
    case '<':
      tok = newToken(token.LT, lex.ch)
    case '>':
      tok = newToken(token.GT, lex.ch)
    case ',':
      tok = newToken(token.COMMA, lex.ch)
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
    case 0:
      tok.Literal = ""
      tok.Type = token.EOF
    default:
      if isLetter(lex.ch) {
        tok.Literal = lex.readIdentifier()
        tok.Type = token.LookupIdentifier(tok.Literal)
        return tok
      } else if isDigit(lex.ch) {
        tok.Literal = lex.readNumber()
        tok.Type = token.INT
        return tok
      } else {
        tok = newToken(token.ILLEGAL, lex.ch)
      }
  }

  lex.readChar()
  return tok
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

func (lex *Lexer) readIdentifier() string {
  var position int = lex.cursor
  for isLetter(lex.ch){
    lex.readChar()
  }

  return lex.input[position:lex.cursor]
}


func (lex *Lexer) readNumber() string{
  position := lex.cursor
  for isDigit(lex.ch) {
    lex.readChar()
  }

  return lex.input[position:lex.cursor]
}


func (lex *Lexer) peekChar() byte{
  if lex.readCursor >= len(lex.input) {
    return 0
  } else {
    return lex.input[lex.readCursor]
  }
}


/************************************* 
  Helper Functions 
*************************************/

// Token Creator
func newToken(tokenType token.TokenType, ch byte) token.Token{
  return token.Token{Type: tokenType, Literal: string(ch)}

}

func isLetter(ch byte) bool {
  return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && '9' >= ch
}

func (lex *Lexer) skipWhiteSpace(){
  for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
    lex.readChar()
  }
}
