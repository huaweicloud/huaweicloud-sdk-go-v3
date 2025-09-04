package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveJobGroupRequestBodyJobs 任务组
type MoveJobGroupRequestBodyJobs struct {

	// 任务编号
	JobId *string `json:"job_id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`
}

func (o MoveJobGroupRequestBodyJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveJobGroupRequestBodyJobs struct{}"
	}

	return strings.Join([]string{"MoveJobGroupRequestBodyJobs", string(data)}, " ")
}
