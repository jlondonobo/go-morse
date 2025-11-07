package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func sayHello(context.Context, *cli.Command) error {
	fmt.Println("Hello ")
	return nil
}

func main() {
	cmd := &cli.Command{
		Name:   "gorse",
		Usage:  "beep beep beeep",
		Action: sayHello,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
