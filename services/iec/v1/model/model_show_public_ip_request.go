package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPublicIpRequest struct {
	// 弹性公网IP ID。

	PublicipId string `json:"publicip_id"`
}

func (o ShowPublicIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicIpRequest", string(data)}, " ")
}
