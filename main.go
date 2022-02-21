package main

import (
  "fmt"
  //"encoding/json"
  "os"
)

// Notes:
// flag for CLI
// https://pkg.go.dev/flag

// define struct for JSON data, may not need
type Incident struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Discovered string `json:"discovered"`
  Description string `json:"description"`
  Status string `json:"status"`
}

func main() {
    // Open the JSON data file for usage
    jsonFile, err := os.Open("input/data.json")

    // if file not found, print error
    if err != nil {
      fmt.Println("Error Accessing File:\n")
      fmt.Println(err)
      fmt.Println("\nEnsure File is in 'input' folder & named 'data.json'.")
    } else {
      // file access successful!
      fmt.Println("File Successfully Accessed")
      // defer closing the file to allow parsing
      defer jsonFile.Close()
    }


    //fmt.Println("Hello World! from Andrew Giardina")
}
