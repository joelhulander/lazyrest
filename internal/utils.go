package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

var file *os.File
var logger *slog.Logger

func Setup() {
	logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	if _, ok := os.LookupEnv("DEBUG"); ok {
		file = newLogFile("app.log")
		logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{ Level: slog.LevelDebug }))
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
		logger.Error("error occurred", "err", err)
		os.Exit(1)
	}
	

	if err = os.MkdirAll(filepath.Join(directory, "lazyrest"), 0755); err != nil {
		logger.Error("Error occurred while creating directories", "err", err)
	}

	return filepath.Join(directory, "lazyrest")
}

func Cleanup() {
	if file != nil {
		file.Close()
	}
}

