package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceRequest struct {
	// 边缘实例ID。

	ServerId string `json:"server_id"`

	Body *UpdateInstanceRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
