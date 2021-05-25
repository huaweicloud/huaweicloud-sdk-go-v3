package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResumeScalingGroupRequest struct {
	// 伸缩组ID

	ScalingGroupId string `json:"scaling_group_id"`

	Body *ResumeScalingGroupOption `json:"body,omitempty"`
}

func (o ResumeScalingGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResumeScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"ResumeScalingGroupRequest", string(data)}, " ")
}
