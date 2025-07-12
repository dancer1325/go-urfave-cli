package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Use a randomized port",
				Value:       0,
				DefaultText: "random",
				Action: func(ctx *cli.Context, v int) error {
					fmt.Println("Flag.Action with value ", v)
					if v >= 65536 {
						return fmt.Errorf("Flag port value %v out of range[0-65535]", v)
					}
					return nil
				},
			},
		},
		// ⚠️App.Action != App.Flags[*].Action⚠️
		Action: func(cCtx *cli.Context) error {
			port := cCtx.Int("port")
			fmt.Printf("Port is: %d\n", port)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
