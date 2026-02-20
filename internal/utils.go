package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

)

var errorFile *os.File
var errorLogger *log.Logger
var messagesFile *os.File
var messagesLogger *log.Logger

func Setup() {
	if _, ok := os.LookupEnv("DEBUG"); ok {
		errorFile = newLogFile("errors.log")
		errorLogger = log.New(errorFile, "ERROR: ", log.LstdFlags)
		messagesFile = newLogFile("messages.log")
		messagesLogger = log.New(messagesFile, "MESSAGE: ", log.LstdFlags)
	}
}

func newLogFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
		os.Exit(1)
	}

	return file

}

func GetRootDir() string {
	if dir, exists := os.LookupEnv("XDG_DATA_HOME"); exists {
		return filepath.Join(dir, "lazyrest")
	} 
	directory, err := os.UserHomeDir()
	if err != nil {
		errorLogger.Println(err)
		os.Exit(1)
	}
	

	if err = os.MkdirAll(filepath.Join(directory, "lazyrest"), 0755); err != nil {
		errorLogger.Println(err)
	}

	return filepath.Join(directory, "lazyrest")
}

func Cleanup() {
	if errorFile != nil {
		errorFile.Close()
	}
	if messagesFile != nil {
		messagesFile.Close()
	}
}

