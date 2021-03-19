package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	AppId string `json:"app_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
