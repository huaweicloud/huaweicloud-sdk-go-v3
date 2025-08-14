package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventName **参数解释**： 事件名称 **取值范围**： 字符长度1-256位
type EventName struct {
}

func (o EventName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventName struct{}"
	}

	return strings.Join([]string{"EventName", string(data)}, " ")
}
