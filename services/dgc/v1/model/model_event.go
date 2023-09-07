package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event struct {

	// 事件类型
	EventType string `json:"eventType"`

	// DIS通道名称
	Channel string `json:"channel"`

	// 执行失败处理策略
	FailPolicy *string `json:"failPolicy,omitempty"`

	// 调度并发数
	Concurrent *int32 `json:"concurrent,omitempty"`

	// 读取策略，LAST ：从上次位置读取，NEW- 从最新位置读取，默认为LAST
	ReadPolicy *string `json:"readPolicy,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
