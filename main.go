package main

import (
  "fmt"
  "encoding/json"
  "encoding/csv"
  "io/ioutil"
  "os"
  "strconv"
  //"flag"
)

// CLI flag toggles
// default settings: Sort ascending, by date discovered (ignore status)
var directionAscending, sortDiscovered, sortStatus bool = true, true, false

// Notes:
// flag for CLI
// https://pkg.go.dev/flag

// define struct for JSON data
type Incident struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Discovered string `json:"discovered"`
  Description string `json:"description"`
  Status string `json:"status"`
}

// define struct for JSON array
type IncidentList struct {
  IncidentList []Incident `json:""` // key/value code may not be necessary
}

func main() {
    // to format first line of csv file
    // todo: improve this, have the actual json key names used
    columnNames := []string {"id","name","discovered","description","status"}

    // Open the JSON data file for usage
    jsonFile, err := os.Open("input/data.json")

    // if file not found, print error
    if err != nil {
      fmt.Println("Error Accessing File:\n")
      fmt.Println(err)
      fmt.Println("\nEnsure File is in 'input' folder & named 'data.json'.")
    } else {
      // file access successful!
      fmt.Println("Input File Successfully Accessed")
    }

    // defer closing the file to allow parsing
    defer jsonFile.Close()
    // read JSON file as byte array
    byteValue, _ := ioutil.ReadAll(jsonFile)
    // initialize IncidentList struct
    var ilist IncidentList
    // unmarshal byteValue array into ilist
    err = json.Unmarshal(byteValue, &ilist.IncidentList)

    // if unmarshal error occurs, print error
    if err != nil {
      fmt.Println("Error Reading JSON File:\n")
      fmt.Println(err)
      fmt.Println("\nEnsure File follows expected JSON format.")
    } else {
      // test printing output
      /*
      for i := 0; i < len(ilist.IncidentList); i++ {
        fmt.Println("id: " + strconv.Itoa( ilist.IncidentList[i].Id ))
        fmt.Println("name: " + ilist.IncidentList[i].Name)
        fmt.Println("discovered: " + ilist.IncidentList[i].Discovered)
        fmt.Println("description: " + ilist.IncidentList[i].Description)
        fmt.Println("status: " + ilist.IncidentList[i].Status)
      }
      */
      
      // create csv file in 'output' folder
      csvFile, err := os.Create("output/data.csv")

      // if file creation error occurs, print error
      if err != nil {
        fmt.Println("Error Creating CSV File:\n")
        fmt.Println(err)
        fmt.Println("\nEnsure ")
      } else {
        // file creation successful

        // defer CSV file from closing
        defer csvFile.Close()
        // create writer to write to output file
        writer := csv.NewWriter(csvFile)

        // write the column names as first line
        writer.Write(columnNames)
        // write all JSON data into output CSV file
        for _, dataEntry := range ilist.IncidentList {
          var csvRow []string
          csvRow = append(csvRow, strconv.Itoa( dataEntry.Id ))
          csvRow = append(csvRow, dataEntry.Name)
          csvRow = append(csvRow, dataEntry.Discovered)
          csvRow = append(csvRow, dataEntry.Description)
          csvRow = append(csvRow, dataEntry.Status)
          writer.Write(csvRow)
        }
        fmt.Println("Output File Successfully Created")
        writer.Flush()
      }
    }



    //fmt.Println("Hello World! from Andrew Giardina")
}
