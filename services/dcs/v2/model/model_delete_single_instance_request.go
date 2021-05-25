package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSingleInstanceRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o DeleteSingleInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSingleInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteSingleInstanceRequest", string(data)}, " ")
}
