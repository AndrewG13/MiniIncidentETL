package main

import (
  "fmt"
  "encoding/json"
  "encoding/csv"
  "io/ioutil"
  "os"
  "strconv"
  "time"
  "flag"
)

// CLI flag toggles
// default settings: Sort ascending, by date discovered (ignore status)
var directionAscending bool = false
var sortDiscovered bool = true
var sortStatus bool = false

// define format of Incident dates
const dateFormat = "2006-01-02" // special date?

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

/*
*  Swap Function
*    swaps two elements in the IncidentList
*/
func swap(list IncidentList, i int, j int) {
  list.IncidentList[i], list.IncidentList[j] = list.IncidentList[j], list.IncidentList[i]
  return
}

/*
*  Sort On Discovered Function
*    Applies sorting based on date discovered
*/
func sortOnDisc(list IncidentList) {
  for i := 0; i < len(list.IncidentList); i++ {
    smallest := i
    for j := i + 1; j < len(list.IncidentList); j++ {
      // compare dates
      if compareDates(list.IncidentList[smallest], list.IncidentList[j]) {
        smallest = j
      }
    }
    // swap current index i with smallest found element
    swap(list, smallest, i)
  }
  return
}

/*
*  Sort On Status Function
*    Applies sorting based on Incident status
*    Algorithm: In-place Selection Sort
*/
func sortOnStat(list IncidentList) {
  for i := 0; i < len(list.IncidentList); i++ {
    // keep track of current smallest index
    smallest := i
    for j := i + 1; j < len(list.IncidentList); j++ {
      // check statuses
      if compareStatus(list.IncidentList[smallest], list.IncidentList[j]) {
        smallest = j
      }
    }
    // swap current index i with smallest found element
    swap(list, smallest, i)
  }
  return
}

/*
*  Compare Status Function
*    helper function used to compare two Incident Status values
*    New < In Progress < Done
*/
func compareStatus(in1, in2 Incident) bool {
  // determine value of status for each Incident
  inVal1 := statusValue(in1)
  inVal2 := statusValue(in2)

  if directionAscending {
    return inVal1 <= inVal2
  } else {
    return inVal1 > inVal2
  }
}

/*
*  Status Value Function
*    returns the value of the Incident's status
*    This is used for comparison purposes
*    New = 3
*    In Progress = 2
*    Done = 1
*/
func statusValue(in Incident) int {
  // return correct Status numeric value
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
*  Compare Status Function
*    helper function used to compare two Incident Dates
*/
func compareDates(in1, in2 Incident) bool {
  // parse dates for each Incident
  inDate1, _ := time.Parse(dateFormat, in1.Discovered)
  inDate2, _ := time.Parse(dateFormat, in2.Discovered)
  // Date comparison
  // https://stackoverflow.com/questions/45024526/comparing-two-dates-without-taking-time-into-account
  oneDay := 24 * time.Hour
	inDate1 = inDate1.Truncate(oneDay)
	inDate2 = inDate2.Truncate(oneDay)
  // numeric result of both dates compared
	before := inDate1.Sub(inDate2)

  if directionAscending {
    // in1 is before in2
    return before > 0
  } else {
    // in1 is after in2
    return before <= 0
  }
}

/*
*  Main
*/
func main() {
    // format for the first line of csv file
    // todo: improve this, have the actual json key names used
    columnNames := []string {"ID", "Name", "Discovered", "Description", "Status"}
    // acceptable sort directions

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

      // check if sorting is necessary (A size of 1 is sorted)
      if len(ilist.IncidentList) > 1 {
        // if user entered sorting preference, call requested sorting mode
        if sortStatus {
          sortOnStat(ilist)
        } else
        if sortDiscovered {
          sortOnDisc(ilist)
        }
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

/*
*  flags & cmd
*/

    // -sortfield
    //  Specify 'Discovered' or 'Status' to sort on
    sortfieldCmd := flag.NewFlagSet("sortfield", flag.ExitOnError)
    sortfieldStat := sortfieldCmd.String("status", "", "status")
    sortfieldDisc := sortfieldCmd.String("discovered", "", "discovered")

    // -sortdirection
    //  Specify 'Ascending' or 'Descending' Direction to sort by
    sortdirectionCmd := flag.NewFlagSet("sortdirection", flag.ExitOnError)
    sortdirectionAs := sortdirectionCmd.String("ascending", "", "ascending")
    sortdirectionDs := sortdirectionCmd.String("descending", "", "descending")

    switch os.Args[1] {
      case "sortfield":

      case "sortdirection":

    }

flag.Parse()

//fmt.Println(sortfieldPtr, sortdirectionPtr)
fmt.Println("\nAscending Mode: ", directionAscending)
fmt.Println("    sortStatus: ", sortStatus)
fmt.Println("sortDiscovered: ", sortDiscovered)

// log program termination
fmt.Println("\nProgram Terminated")
}

func handleSortField(sortfieldCmd) {

}

func handleSortDirection(sortdirectionCmd *flag.FlagSet, status *string, discovered *string) {

}
