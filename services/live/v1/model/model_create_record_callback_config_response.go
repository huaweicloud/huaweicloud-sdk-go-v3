package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRecordCallbackConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRecordCallbackConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRecordCallbackConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordCallbackConfigResponse", string(data)}, " ")
}
