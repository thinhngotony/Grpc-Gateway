package proto

import (
	"context"
	"main/proto/protogen"
	"main/utils"

	"github.com/Azure/go-autorest/autorest/to"
	log "github.com/sirupsen/logrus"
)

type YourServiceServer struct{}

func (YourServiceServer) Sample(ctx context.Context, req *protogen.Interface) (*protogen.Interface, error) {
	log.Infof("[Sample] Received message: \n%v", utils.InterfaceToString(req))
	return &protogen.Interface{StrVal: to.StringPtr("Hello world")}, nil
}
