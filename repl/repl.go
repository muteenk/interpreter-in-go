package repl

import (
  "fmt"
  "bufio"
  "io"
  "monkey/lexer"
  "monkey/token"
)


func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(">> ")
    scanned := scanner.Scan()
    
    if !scanned {
      return
    }

    line := scanner.Text()
    lex := lexer.NewLexer(line)
  
    for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
      fmt.Printf("%+v\n", tok)
    }

  }
}


