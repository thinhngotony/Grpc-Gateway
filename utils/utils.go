package utils

import (
	"encoding/json"
	"fmt"
)

func InterfaceToString(data interface{}) string {
	manifestJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", data)
	}
	return string(manifestJson)
}
