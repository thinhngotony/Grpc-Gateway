package proto

import (
	"main/proto/protogen"
	"main/utils"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/sirupsen/logrus"
)

func GetVersion() *protogen.BuildInfo {
	bi := protogen.BuildInfo{
		ApiVersion: to.StringPtr(utils.ApiVersion),
		Commit:     to.StringPtr(utils.Commit),
		BuildDate:  to.StringPtr(utils.BuildDate),
	}

	logrus.Infof("Version: %v, Commit: %v, Build Date: %v\n\n", bi.ApiVersion, bi.Commit, bi.BuildDate)
	return &bi
}
