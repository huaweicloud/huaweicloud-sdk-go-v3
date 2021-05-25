package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteConnctionRequest struct {
	// 连接名称.

	ConnectionName string `json:"connection_name"`
}

func (o DeleteConnctionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteConnctionRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnctionRequest", string(data)}, " ")
}
