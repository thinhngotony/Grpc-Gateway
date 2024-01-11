package proto

import (
	"context"
	"main/proto/protogen"
	"main/utils"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/anypb"
)

type YourServiceServer struct{}

func (YourServiceServer) Sample(ctx context.Context, req *protogen.Interface) (*protogen.Interface, error) {
	logrus.Infof("[Sample] Received message: \n%v", utils.InterfaceToString(req))
	return &protogen.Interface{Data: &protogen.Interface_StrVal{StrVal: "Hello world"}}, nil
}

func (YourServiceServer) GetVersion(ctx context.Context, req *protogen.Interface) (*protogen.Interface, error) {
	logrus.Infof("[GetVersion] Received request")

	resp, err := anypb.New(GetVersion())
	if err != nil {
		return nil, err
	}

	return &protogen.Interface{Data: &protogen.Interface_AnyVal{AnyVal: resp}}, nil
}
