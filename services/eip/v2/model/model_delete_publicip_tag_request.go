package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePublicipTagRequest struct {
	// 弹性公网IP的id

	PublicipId string `json:"publicip_id"`
	// 标签的键

	Key string `json:"key"`
}

func (o DeletePublicipTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicipTagRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicipTagRequest", string(data)}, " ")
}
