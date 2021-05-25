package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPortRequest struct {
	// 端口ID

	PortId string `json:"port_id"`
}

func (o ShowPortRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPortRequest struct{}"
	}

	return strings.Join([]string{"ShowPortRequest", string(data)}, " ")
}
