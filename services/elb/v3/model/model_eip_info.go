package model

import (
	"encoding/json"

	"strings"
)

// 弹性ip，同publicips
type EipInfo struct {
	// eip_id

	EipId *string `json:"eip_id,omitempty"`
	// eip_address

	EipAddress *string `json:"eip_address,omitempty"`
}

func (o EipInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
