package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteFlinkJobSavepointRequest Request Object
type ExecuteFlinkJobSavepointRequest struct {

	// 正在运行的Flink作业的作业ID。
	JobId string `json:"job_id"`

	Body *ExecuteFlinkJobSavepointRequestBody `json:"body,omitempty"`
}

func (o ExecuteFlinkJobSavepointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteFlinkJobSavepointRequest struct{}"
	}

	return strings.Join([]string{"ExecuteFlinkJobSavepointRequest", string(data)}, " ")
}
