package model

import (
	"encoding/json"

	"strings"
)

// network_info
type UpdateTaskStatusRequestBodyNetworkInfo struct {
	// network_type

	NetworkType string `json:"network_type"`
}

func (o UpdateTaskStatusRequestBodyNetworkInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusRequestBodyNetworkInfo struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusRequestBodyNetworkInfo", string(data)}, " ")
}
