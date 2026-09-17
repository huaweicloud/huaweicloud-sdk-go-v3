package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKillTaskRequest Request Object
type StopKillTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *StopKillTaskRequestBody `json:"body,omitempty"`
}

func (o StopKillTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKillTaskRequest struct{}"
	}

	return strings.Join([]string{"StopKillTaskRequest", string(data)}, " ")
}
