package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateInstanceResponse struct {
	// 操作记录id

	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}
