
# Problem 1: Boss Baby's Revenge

## TechStack
All are written in **Go**,

## How to run

To get started, follow these steps:

1. **Clone the Repository:**
   ```
   git clone https://github.com/BoomNooB/boss-babys-revenge.git
   cd boss-babys-revenge
   ```
   

4. **Running a makefile:**\
   Simply just type this to your terminal after `cd` from step one
   `make run`
   then it will do a couple of thing here
   - `go mod tidy` for download and install dependencies  
   - `go run ./cmd/app/main.go` for running the program
  
	After that it will running with some testcase in `constant.go` file 
	and you can add your own test case by creating new variable in  `constant.go`
   `TCSN = "SOMETESTCASE"`  	
	and then add `constant.TCSN  ` into `main.go` file in `testCases` variable
	
    You can generate test case with 
    `tools.TestCaseGenGoodBoy(numberOfString,isGoodBoy,isThisCaseShouldInvalid)`  
    It will write text to file name `"testCase"`
    Then you can simply copy and paste to `constant.go` 
    And assign to new variable  
    Also dont fotget to add test case in `main.go` file
    This generate test case was comment in `main.go` line `37-41`

   And you can also run `make test` for running unit test file with verbose

---
## What can be improve
1. should pick just one `true` or `false` as `goodBoy`, current code might confused person who try to read code