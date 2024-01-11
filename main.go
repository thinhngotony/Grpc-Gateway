package main

import (
	"flag"
	"main/proto"
	"main/route"

	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	proto.GetVersion()

	go func() {
		if err := route.RunProxy(); err != nil {
			logrus.Errorf("Failed to serve Proxy server %v", err)
			return
		}
	}()

	if err := route.RunGrpc(); err != nil {
		logrus.Errorf("Failed to serve GRPC server %v", err)
		return
	}

}
