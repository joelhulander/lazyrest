package main

import (
	"fmt"
	"os"

	"codeberg.org/joelhulander/lazyrest/internal"
)

func main() {
	app := internal.NewApp()

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
