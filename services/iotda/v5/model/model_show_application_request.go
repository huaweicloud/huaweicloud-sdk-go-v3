package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApplicationRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	AppId      string  `json:"app_id"`
}

func (o ShowApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationRequest", string(data)}, " ")
}
