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
var directionAscending, sortDiscovered, sortStatus bool = true, false, true

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

// function for testing purposes
func change(list IncidentList) {
  for i := 0; i < len(list.IncidentList); i++ {
    if directionAscending {
      list.IncidentList[i].Id = 100
    } else {
      list.IncidentList[i].Id = 1000
    }
  }
  return
}

func swap(list IncidentList, i int, j int) {
  list.IncidentList[i], list.IncidentList[j] = list.IncidentList[j], list.IncidentList[i]
  return
}

func sortOnDisc(list IncidentList) {
  for i := 0; i < len(list.IncidentList); i++ {
    list.IncidentList[i].Id = 100
  }
  return
}

func sortOnStat(list IncidentList) {
  for i := 0; i < len(list.IncidentList); i++ {
    // keep track of current smallest index
    smallest := i
    for j := i + 1; j < len(list.IncidentList); j++ {
      // check if
      if compareStatus(list.IncidentList[smallest], list.IncidentList[j]) {
        smallest = j
      }
    }
    swap(list, smallest, i)
  }
  return
}

func compareStatus(in1, in2 Incident) bool {
  inVal1 := statusValue(in1)
  inVal2 := statusValue(in2)
  if directionAscending {
    return inVal1 <= inVal2
  } else {
    return inVal1 > inVal2
  }
}

func statusValue(in Incident) int {
  var inVal int
  if in.Status == "New" {
      inVal = 3
  } else if in.Status == "In Progress" {
      inVal = 2
  } else {
      inVal = 1
  }
  return inVal
}

/*
*  Main
*/
func main() {
    // format for the first line of csv file
    // todo: improve this, have the actual json key names used
    columnNames := []string {"ID", "Name", "Discovered", "Description", "Status"}
    // acceptable sort directions
    //sortDirections := []string {"Ascending","Descending"}

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

      // test change function
      if len(ilist.IncidentList) > 1 {
        sortOnStat(ilist)
      }

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
  }


    fmt.Println("Program Terminated")
}
