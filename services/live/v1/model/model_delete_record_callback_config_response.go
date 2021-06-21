package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRecordCallbackConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRecordCallbackConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordCallbackConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordCallbackConfigResponse", string(data)}, " ")
}
