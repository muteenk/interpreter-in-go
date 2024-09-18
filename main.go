package main

import (
  "fmt"
  "os"
  "monkey/repl"
);

func main() {
  fmt.Println("Monkey Language Interpreter 1.0.0\nBuilt from Scratch in Go Lang\nBy Muteen K.\n\n Feel free to type in Commands\n")
  
  repl.Start(os.Stdin, os.Stdout)
}
