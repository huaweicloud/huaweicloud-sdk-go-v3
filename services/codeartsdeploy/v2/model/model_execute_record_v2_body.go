package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteRecordV2Body 应用部署记录Body
type ExecuteRecordV2Body struct {

	// 部署用时
	Duration *string `json:"duration,omitempty"`

	// 应用状态
	State *string `json:"state,omitempty"`

	// 操作人用户名
	Operator *string `json:"operator,omitempty"`

	// 部署记录id
	ExecutionId *string `json:"execution_id,omitempty"`

	// 部署应用的开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 操作人昵称
	Nickname *string `json:"nickname,omitempty"`

	// 部署应用的结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 部署记录序列号
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
