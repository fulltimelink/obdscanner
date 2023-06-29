package main

import (
	"fmt"
	"github.com/rzetterberg/elmobd"
)

func main() {

	dev, err := elmobd.NewDevice("/COM4", false)

	if err != nil {
		fmt.Println("Failed to create new device", err)
		return
	}

	supported, err := dev.CheckSupportedCommands()

	if err != nil {
		fmt.Println("Failed to check supported commands", err)
		return
	}

	allCommands := elmobd.GetSensorCommands()
	carCommands := supported.FilterSupported(allCommands)

	fmt.Printf("%d of %d commands supported:\n", len(carCommands), len(allCommands))

	for _, cmd := range carCommands {
		fmt.Printf("- %s supported\n", cmd.Key())
	}

}
