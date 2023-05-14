package main

import (
	"fmt"
	"log"

	projector_config "github.com/zolwiastyl/submarine/pkg/config"
	"github.com/zolwiastyl/submarine/pkg/projector"
)

func main() {
	opts, err := projector_config.GetOps()
	if err != nil {
		log.Fatal("Terribly wrong input")
	}
	config, err := projector_config.NewConfig(opts)
	if err != nil {
		log.Fatalf("Unable to get options %v", err)
	}

	fmt.Printf("config: %+v", config)
	fmt.Printf("opts: %+v", opts)

	projector := projector.NewProjector(config)

	switch config.Operation {
	case projector_config.Print:
		fmt.Println("print")

		if config.Args == nil || len(config.Args) == 0 {
			projector.GetValueAll()
			break
		}
		projector.GetValue(config.Args[0])

	case projector_config.Add:
		fmt.Println("add")

		projector.SetValue(config.Args[0], config.Args[1])
		projector.Save()

	case projector_config.Remove:
		fmt.Println("remove")

		projector.RemoveValue(config.Args[0])
		projector.Save()
	}
}
