package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteInstanceResponse struct {
	// 返回值

	Result *bool `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}
