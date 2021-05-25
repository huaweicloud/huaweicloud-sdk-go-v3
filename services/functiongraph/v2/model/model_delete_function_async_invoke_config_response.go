package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFunctionAsyncInvokeConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionAsyncInvokeConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
