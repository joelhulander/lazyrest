package main

import (
	"fmt"
	"os"

	"github.com/joelhulander/lazyrest/internal"
)


func main() {
	internal.Setup()
	defer internal.Cleanup()
	app := internal.NewApp(internal.GetRootDir())

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

