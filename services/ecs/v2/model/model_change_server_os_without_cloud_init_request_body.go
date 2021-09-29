package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type ChangeServerOsWithoutCloudInitRequestBody struct {
	OsChange *ChangeServerOsWithoutCloudInitOption `json:"os-change"`
}

func (o ChangeServerOsWithoutCloudInitRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitRequestBody", string(data)}, " ")
}
