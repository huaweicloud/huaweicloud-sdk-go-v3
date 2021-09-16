package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePublicIpRequest struct {
	// 弹性公网IP ID

	PublicipId string `json:"publicip_id"`
}

func (o DeletePublicIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicIpRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicIpRequest", string(data)}, " ")
}
