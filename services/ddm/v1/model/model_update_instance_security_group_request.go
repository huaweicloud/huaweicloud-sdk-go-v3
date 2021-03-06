package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceSecurityGroupRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`

	Body *ModifyInstanceSecurityGroupReq `json:"body,omitempty"`
}

func (o UpdateInstanceSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceSecurityGroupRequest", string(data)}, " ")
}
