package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestartInstanceRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`

	Body *RestartInstanceReq `json:"body,omitempty"`
}

func (o RestartInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestartInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequest", string(data)}, " ")
}
