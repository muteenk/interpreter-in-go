package main

import (
  "fmt"
)


func main() {
  fmt.Println("Hello! This is the Monkey programming language!")
  fmt.Println("Feel free to type in commands")
  var name string;
  fmt.Scan(&name);

  fmt.Printf("Hello %s\n", name);
}
