package main

import (
	"flag"
	"main/route"

	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	go func() {
		if err := route.RunProxy(); err != nil {
			log.Errorf("Failed to serve Proxy server %v", err)
			return
		}
	}()

	if err := route.RunGrpc(); err != nil {
		log.Errorf("Failed to serve GRPC server %v", err)
		return
	}

}
