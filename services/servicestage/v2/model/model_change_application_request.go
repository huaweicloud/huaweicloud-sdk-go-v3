package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeApplicationRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`

	Body *ApplicationModify `json:"body,omitempty"`
}

func (o ChangeApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeApplicationRequest struct{}"
	}

	return strings.Join([]string{"ChangeApplicationRequest", string(data)}, " ")
}
