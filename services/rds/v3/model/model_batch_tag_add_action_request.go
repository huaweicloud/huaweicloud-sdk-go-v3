package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchTagAddActionRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *BatchTagActionAddRequestBody `json:"body,omitempty"`
}

func (o BatchTagAddActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTagAddActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionRequest", string(data)}, " ")
}
