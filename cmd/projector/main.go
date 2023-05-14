package main

import (
	"encoding/json"
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

	projector := projector.NewProjector(config)

	switch config.Operation {
	case projector_config.Print:

		if len(config.Args) == 0 {
			result := projector.GetValueAll()
			marshaled, err := json.Marshal(result)
			if err != nil {
				log.Fatalf("Unable to marshal %v", err)
			}
			fmt.Printf("%v", string(marshaled))
			break
		} else if result, ok := projector.GetValue(config.Args[0]); ok {
			fmt.Printf("%v", result)
			break
		}

	case projector_config.Add:
		projector.SetValue(config.Args[0], config.Args[1])
		projector.Save()

	case projector_config.Remove:

		projector.RemoveValue(config.Args[0])
		projector.Save()
	}
}
