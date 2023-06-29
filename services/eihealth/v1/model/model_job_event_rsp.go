package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEventRsp 作业事件响应内容
type JobEventRsp struct {

	// 执行动作名称
	ActionName *string `json:"action_name,omitempty"`

	// 作业启动事件发生次数
	Count *int32 `json:"count,omitempty"`

	// 作业启动事件首次上报时间
	FirstTimestamp *string `json:"first_timestamp,omitempty"`

	// 作业启动事件末次上报时间
	LastTimestamp *string `json:"last_timestamp,omitempty"`

	// 作业启动事件详细信息
	Message *string `json:"message,omitempty"`

	// 作业启动事件状态
	Reason *string `json:"reason,omitempty"`

	// 作业启动事件类型
	Type *string `json:"type,omitempty"`
}

func (o JobEventRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEventRsp struct{}"
	}

	return strings.Join([]string{"JobEventRsp", string(data)}, " ")
}
