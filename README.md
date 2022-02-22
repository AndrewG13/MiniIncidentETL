# MiniIncidentETL
RadarFirst Interview Exercise - Andrew Giardina
  
MiniIncidentETL is a CLI tool created in Go that will convert an appropriate JSON file into CSV format.   

## How to Operate:
- Simply create an appropriately formatted JSON file (or use one of my test samples in the 'tests' directory)
- Name the file: **data.json**
- Insert the file into the 'input' directory
- Open a terminal in the main directory
- Type in & run: `go run main.go`, applying any additional sorting preferences desired (see **Commands & Examples**)
- Check out your new CSV file located in the 'output' directory

## How to Install:
- Clone this repository:  
`git clone https://github.com/AndrewG13/MiniIncidentETL.git`  

## Commands & Examples:  
Available Commands -  
- sortfield <field>  [status, discovered]  
  Sorts Incidents by field specified  
- sortdirection <direction> [ascending, descending]  
  Sorts Incidents in order specified  
- columns <attributes,...> [id, name, discovered, description, status]  
  Filters which attributes to include in the CSV file
## Debug Mode:  
Enables additional print statements to verify program behaviour
- Line 20: `var debug bool = false`
  Enable by changing to `true`  
  
Helpful Examples - 
- `go run main.go`
- `go run main.go sortfield -status`  
- `go run main.go sortfield -discovered` sortdirection -ascending  
- `go run main.go columns -id -status -name sortdirection -descending sortfield -status`
  
Error Examples (for testing purposes) -
- Invalid tag: `go run main.go sortfield -nah`  
- Duplicate Command: `go run main.go sortfield -status sortfield -status`
- Invalid Command: `go run main.go sortfield -discovered randomstring columns -id`

## Future Development Plans:  
- Improve the sorting algorithms used for sortfield (Discovered: Line 92 & Status: Line 112)
- Improve duplicate command verification (it's repetitive, starting at Line 284-307)
- Implement the command `-notdone`, to filter only Incidents not done/completed  
- Implement the command `-alldone`, to filter onlg Incidents done/completed  
