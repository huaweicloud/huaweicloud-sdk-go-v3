package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddScalingInstancesRequest struct {
	// 实例ID。

	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchAddInstancesOption `json:"body,omitempty"`
}

func (o BatchAddScalingInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchAddScalingInstancesRequest", string(data)}, " ")
}
