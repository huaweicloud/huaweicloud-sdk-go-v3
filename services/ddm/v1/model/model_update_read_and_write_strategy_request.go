package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateReadAndWriteStrategyRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`

	Body *ModifyReadAndWriteStrategyReq `json:"body,omitempty"`
}

func (o UpdateReadAndWriteStrategyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReadAndWriteStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateReadAndWriteStrategyRequest", string(data)}, " ")
}
