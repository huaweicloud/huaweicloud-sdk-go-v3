package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateClientNetworkRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ClientNetworkRequestBody `json:"body,omitempty"`
}

func (o UpdateClientNetworkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkRequest struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkRequest", string(data)}, " ")
}
