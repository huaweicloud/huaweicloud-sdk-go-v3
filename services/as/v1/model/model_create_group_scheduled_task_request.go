package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupScheduledTaskRequest Request Object
type CreateGroupScheduledTaskRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	Body *CreateScheduledTaskOption `json:"body,omitempty"`
}

func (o CreateGroupScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupScheduledTaskRequest", string(data)}, " ")
}
