package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowScalingGroupRequest struct {
	// 伸缩组ID。

	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowScalingGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingGroupRequest", string(data)}, " ")
}
