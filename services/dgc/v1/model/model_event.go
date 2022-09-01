package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event struct {

	// 事件类型
	EventType *string `json:"eventType,omitempty" xml:"eventType"`

	// DIS通道名称
	Channel *string `json:"channel,omitempty" xml:"channel"`

	// 执行失败处理策略
	FailPolicy *string `json:"failPolicy,omitempty" xml:"failPolicy"`

	// 调度并发数
	Concurrent *int32 `json:"concurrent,omitempty" xml:"concurrent"`

	// 读取策略
	ReadPolicy *string `json:"readPolicy,omitempty" xml:"readPolicy"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
