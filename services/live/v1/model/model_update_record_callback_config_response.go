package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateRecordCallbackConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRecordCallbackConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRecordCallbackConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecordCallbackConfigResponse", string(data)}, " ")
}
