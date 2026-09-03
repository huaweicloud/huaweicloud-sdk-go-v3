package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlKillingTaskResp 自动kill会话任务
type SqlKillingTaskResp struct {

	// 用户名
	User *string `json:"user,omitempty"`

	// host地址
	Host *string `json:"host,omitempty"`

	// 数据库
	Db *string `json:"db,omitempty"`

	// 信息
	Info *string `json:"info,omitempty"`

	// 命令行
	Command *string `json:"command,omitempty"`

	// 次数
	Time *int64 `json:"time,omitempty"`

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 任务耗时
	TaskDuration *int64 `json:"task_duration,omitempty"`

	// 任务状态
	TaskState *int32 `json:"task_state,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o SqlKillingTaskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlKillingTaskResp struct{}"
	}

	return strings.Join([]string{"SqlKillingTaskResp", string(data)}, " ")
}
