package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopInstanceResponse struct {
	// 返回值
	Result *bool `json:"result,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopInstanceResponse", string(data)}, " ")
}
