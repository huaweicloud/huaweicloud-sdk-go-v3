package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AsyncInvokeReservedFunctionResponse struct {
	// 预留实例id

	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AsyncInvokeReservedFunctionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AsyncInvokeReservedFunctionResponse struct{}"
	}

	return strings.Join([]string{"AsyncInvokeReservedFunctionResponse", string(data)}, " ")
}
