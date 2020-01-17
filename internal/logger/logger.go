package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Export the loggers for use externally
var GeneralLogger *log.Logger
var ErrorLogger *log.Logger
var TestLogger *log.Logger

func init() {
	// configure directory for logs
	// all new log files must be added under
	// see generalLog example for configuring a log file
	absPath, err := filepath.Abs("../logs")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	// configure general log paths
	generalLog, err := os.OpenFile(absPath+"/general-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// configure log path for tests
	testLog, err := os.OpenFile(absPath+"/test-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// configure logger types using the path, prefix and log format
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	TestLogger = log.New(testLog, "Test Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
