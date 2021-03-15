package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApplicationsRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	DefaultApp *bool   `json:"default_app,omitempty"`
}

func (o ShowApplicationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationsRequest", string(data)}, " ")
}
