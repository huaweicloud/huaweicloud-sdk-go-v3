package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryRecord1 struct {

	// 构建记录id--唯一key
	RecordId *string `json:"record_id,omitempty"`

	// 任务id
	JobId *string `json:"job_id,omitempty"`

	// 构建编号
	BuildNumber *int32 `json:"build_number,omitempty"`

	// 构建开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 构建结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 构建结果
	Result *string `json:"result,omitempty"`

	// 代码分支
	Branch *string `json:"branch,omitempty"`

	// 代码提交的commit id
	CommitId *string `json:"commit_id,omitempty"`

	// 代码提交时用户输入的提交信息，只有使用codehub仓库时有值
	CommitMessage *string `json:"commit_message,omitempty"`

	// 执行构建任务的用户的用户名
	Executor *string `json:"executor,omitempty"`

	// 触发方式，可选值：手工触发，定时触发，代码更新触发，流水线触发
	TriggerType *string `json:"trigger_type,omitempty"`
}

func (o HistoryRecord1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryRecord1 struct{}"
	}

	return strings.Join([]string{"HistoryRecord1", string(data)}, " ")
}
