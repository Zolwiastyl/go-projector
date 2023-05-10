package main

import (
	"fmt"
	"log"

	projector_config "github.com/zolwiastyl/submarine/pkg/config"
)

func main() {
	opts, err := projector_config.GetOps()
	if err != nil {
		log.Fatal("Da fuck man you trying to put into me ")
	}
	config, err := projector_config.NewConfig(opts)
	if err != nil {
		log.Fatalf("Unable to get options %v", err)
	}

	fmt.Printf("config: %+v", config)
	fmt.Printf("opts: %+v", opts)
}
