package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveEvent struct {

	// 事件戳。从1970-01-01 00:00:00:000开始的毫秒数
	Timestamp int64 `json:"timestamp"`

	// 事件类型。
	Type *int32 `json:"type,omitempty"`

	// 事件内容。
	Content *string `json:"content,omitempty"`
}

func (o LiveEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveEvent struct{}"
	}

	return strings.Join([]string{"LiveEvent", string(data)}, " ")
}
