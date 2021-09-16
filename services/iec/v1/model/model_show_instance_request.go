package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceRequest struct {
	// 边缘实例ID。

	ServerId string `json:"server_id"`
}

func (o ShowInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRequest", string(data)}, " ")
}
