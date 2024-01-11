package utils

import (
	"encoding/json"
	"fmt"
)

var (
	ApiVersion string
	Branch     string
	Commit     string
	BuildDate  string
	GoVersion  string
	OsArch     string
)

func InterfaceToString(data interface{}) string {
	manifestJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", data)
	}
	return string(manifestJson)
}
