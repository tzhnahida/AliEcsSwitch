package main

import (
	"flag"
	"log"

	api "github.com/alibabacloud-go/ecs-20140526/v7/client"

	"aliecs"
)

var (
	cfgPath string
	start   bool
	stop    bool
)

func init() {
	flag.StringVar(&cfgPath, "cfg", "config.toml", "set config path")
	flag.BoolVar(&start, "start", false, "start instance")
	flag.BoolVar(&stop, "stop", false, "stop instance")
	flag.Parse()

	if start && stop {
		log.Fatal("cannot use both -start and -stop at the same time")
	}
	if !start && !stop {
		log.Fatal("please use either -start or -stop flag")
	}
}

func main() {
	config, err := aliecs.LoadConfig(cfgPath)
	if err != nil {
		log.Fatal("configuration file reading error")
	}
	client, err := aliecs.CreateClient(config)
	if err != nil {
		log.Fatal("failed to instantiate a client")
	}
	switch {
	case start:
		StartInstance(config, client)
	case stop:
		StopInstance(config, client)
	}
}

func StartInstance(config *aliecs.Config, client *api.Client) {
	log.Println("start the ECS instance")
	err := aliecs.StartInstance(config, client)
	if err != nil {
		log.Fatal("failed to start ECS")
	}
}

func StopInstance(config *aliecs.Config, client *api.Client) {
	log.Println("stop the ECS instance")
	err := aliecs.StopInstance(config, client)
	if err != nil {
		log.Fatal("failed to stop ECS")
	}
}
