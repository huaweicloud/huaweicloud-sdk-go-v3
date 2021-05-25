package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePublicipRequest struct {
	// 弹性公网IP唯一标识

	PublicipId string `json:"publicip_id"`
}

func (o DeletePublicipRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicipRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicipRequest", string(data)}, " ")
}
