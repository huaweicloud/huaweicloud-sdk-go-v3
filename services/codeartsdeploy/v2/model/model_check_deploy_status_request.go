package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDeployStatusRequest Request Object
type CheckDeployStatusRequest struct {

	// task_id
	TaskId string `json:"task_id"`

	// record_id
	RecordId *string `json:"record_id,omitempty"`

	// 是否返回部署任务各步骤的状态， true返回， false不返回
	StepState *bool `json:"step_state,omitempty"`
}

func (o CheckDeployStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDeployStatusRequest struct{}"
	}

	return strings.Join([]string{"CheckDeployStatusRequest", string(data)}, " ")
}
