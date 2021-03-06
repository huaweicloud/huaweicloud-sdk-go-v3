package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {
	Instance *InstancesVo `json:"instance,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
