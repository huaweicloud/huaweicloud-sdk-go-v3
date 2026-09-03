package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventEventsDto binlog事件概览信息
type EventEventsDto struct {

	// 文件名称
	LogName *string `json:"log_name,omitempty"`

	// 位置
	Pos *int64 `json:"pos,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 服务器ID
	ServerId *string `json:"server_id,omitempty"`

	// 结束位置
	EndLogPos *int64 `json:"end_log_pos,omitempty"`

	// 信息
	Info *string `json:"info,omitempty"`
}

func (o EventEventsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventEventsDto struct{}"
	}

	return strings.Join([]string{"EventEventsDto", string(data)}, " ")
}
