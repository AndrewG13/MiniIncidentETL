# MiniIncidentETL
RadarFirst Interview Exercise - Andrew Giardina

## How to Install:
- Clone this repository:  
`git clone https://github.com/AndrewG13/MiniIncidentETL.git`  

## How to Run:  
- Put JSON data in the 'input' directory  
- File must be named 'data.json'  
- Run the program with the following command:  
`go run main.go`  
- The CSV file will be created in the 'output' directory  

## How to Change Fields & Sorting Attributes:  
In the 'main.go' file, the you may alter the following lines as desired:  
- Line 15: Debug Mode  
- Line 19: Sort Ascending  
- Line 20: Sort by Discovered date   
- Line 21: Sort by Incident Status  

## Future Development Plans:  
- Finish up the Command Line Interface commands & arguments for user friendliness  
- Implement the command `-notdone`, to filter only Incidents not done/completed  
- Implement the command `-alldone`, to filter onlg Incidents done/completed  
