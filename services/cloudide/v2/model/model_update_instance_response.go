package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateInstanceResponse struct {
	// 返回值

	Result *bool `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}
