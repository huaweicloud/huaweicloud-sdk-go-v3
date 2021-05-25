package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchUnsetScalingInstancesStandbyRequest struct {
	// 实例ID。

	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchExitStandByInstancesOption `json:"body,omitempty"`
}

func (o BatchUnsetScalingInstancesStandbyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUnsetScalingInstancesStandbyRequest struct{}"
	}

	return strings.Join([]string{"BatchUnsetScalingInstancesStandbyRequest", string(data)}, " ")
}
