package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProductConfigRequest struct {
	// 协议类型

	ProtocolType string `json:"protocol_type"`
}

func (o ShowProductConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProductConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowProductConfigRequest", string(data)}, " ")
}
