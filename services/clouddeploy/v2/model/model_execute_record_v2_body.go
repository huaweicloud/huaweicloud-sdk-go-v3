package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务执行记录Body
type ExecuteRecordV2Body struct {

	// 执行用时
	Duration *string `json:"duration,omitempty"`

	// 任务状态
	State *string `json:"state,omitempty"`

	// 操作人用户名
	Operator *string `json:"operator,omitempty"`

	// 执行记录ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 执行任务的开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 操作人昵称
	Nickname *string `json:"nickname,omitempty"`

	// 执行任务的结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 执行序列号
	ReleaseId *int64 `json:"release_id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o ExecuteRecordV2Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRecordV2Body struct{}"
	}

	return strings.Join([]string{"ExecuteRecordV2Body", string(data)}, " ")
}
