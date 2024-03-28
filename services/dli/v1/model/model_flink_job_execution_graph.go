package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobExecutionGraph 查询作业执行计划。
type FlinkJobExecutionGraph struct {

	// flink作业id。
	Jid *string `json:"jid,omitempty"`

	// flink作业名字。
	Name *string `json:"name,omitempty"`

	// 是否可停止。
	IsStoppable *bool `json:"isStoppable,omitempty"`

	// 作业运行状态。
	State *string `json:"state,omitempty"`

	// 作业启动时间。
	StartTime *int64 `json:"start-time,omitempty"`

	// 作业停止时间。
	EndTime *int64 `json:"end-time,omitempty"`

	// 作业运行时长。
	Duration *int64 `json:"duration,omitempty"`
}

func (o FlinkJobExecutionGraph) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobExecutionGraph struct{}"
	}

	return strings.Join([]string{"FlinkJobExecutionGraph", string(data)}, " ")
}
