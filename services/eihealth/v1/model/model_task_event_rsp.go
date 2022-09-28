package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务启动事件响应内容
type TaskEventRsp struct {

	// 任务启动事件类型
	Type *string `json:"type,omitempty"`

	// 任务启动事件发生次数
	Count *int32 `json:"count,omitempty"`

	// 任务启动事件状态
	Reason *string `json:"reason,omitempty"`

	// 任务启动事件详细信息
	Message *string `json:"message,omitempty"`

	// 任务启动事件首次上报时间
	FirstTimestamp *string `json:"first_timestamp,omitempty"`

	// 任务启动事件末次上报时间
	LastTimestamp *string `json:"last_timestamp,omitempty"`
}

func (o TaskEventRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskEventRsp struct{}"
	}

	return strings.Join([]string{"TaskEventRsp", string(data)}, " ")
}
