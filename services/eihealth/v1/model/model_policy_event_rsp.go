package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyEventRsp struct {

	// 事件产生时间
	Time string `json:"time"`

	// 事件名称
	Name string `json:"name"`

	// 事件类型
	EventType string `json:"event_type"`
}

func (o PolicyEventRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEventRsp struct{}"
	}

	return strings.Join([]string{"PolicyEventRsp", string(data)}, " ")
}
