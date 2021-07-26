package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateHostRequest struct {
	// 域名Id

	InstanceId string `json:"instance_id"`

	Body *UpdateHostRequestBody `json:"body,omitempty"`
}

func (o UpdateHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostRequest", string(data)}, " ")
}
