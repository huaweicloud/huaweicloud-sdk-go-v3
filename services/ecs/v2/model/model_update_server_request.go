package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateServerRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`

	Body *UpdateServerRequestBody `json:"body,omitempty"`
}

func (o UpdateServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerRequest", string(data)}, " ")
}
