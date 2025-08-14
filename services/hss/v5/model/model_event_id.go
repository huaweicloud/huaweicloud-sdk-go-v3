package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventId **参数解释**： 事件ID **取值范围**： 字符长度1-64位
type EventId struct {
}

func (o EventId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventId struct{}"
	}

	return strings.Join([]string{"EventId", string(data)}, " ")
}
