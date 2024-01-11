package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueUser 工作队列用户
type WorkloadQueueUser struct {

	// 用户名
	UserName string `json:"user_name"`

	// 执行计划阶段
	OccupyResourceList []OccupyResource `json:"occupy_resource_list"`

	// 执行结果。
	ExecResult *int32 `json:"exec_result,omitempty"`

	// 执行日志。
	ExecLog *string `json:"exec_log,omitempty"`
}

func (o WorkloadQueueUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueUser struct{}"
	}

	return strings.Join([]string{"WorkloadQueueUser", string(data)}, " ")
}
