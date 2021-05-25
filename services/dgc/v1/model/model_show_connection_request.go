package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowConnectionRequest struct {
	// 连接名称.

	ConnectionName string `json:"connection_name"`
}

func (o ShowConnectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionRequest", string(data)}, " ")
}
