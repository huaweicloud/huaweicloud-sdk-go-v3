package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRecordCallbackConfigRequest struct {
	Body *RecordCallbackConfigRequest `json:"body,omitempty"`
}

func (o CreateRecordCallbackConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordCallbackConfigRequest", string(data)}, " ")
}
