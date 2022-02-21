package main

import (
  // add more
  "fmt"
)

// Notes:
// flag for CLI
// https://pkg.go.dev/flag

// define struct for JSON data, may not need
type incident struct {
  Id int
  Name string
  Discovered string
  Description string
  Status string
}

func main() {
    fmt.Println("Hello World! from Andrew Giardina")
}
