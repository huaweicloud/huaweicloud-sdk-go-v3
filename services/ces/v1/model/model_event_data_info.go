package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDataInfo
type EventDataInfo struct {

	// **参数解释**： 事件类型 **取值范围**： 不涉及
	Type string `json:"type"`

	// **参数解释**： 上报时间，UNIX时间戳，单位毫秒 **取值范围**： 不涉及
	Timestamp int64 `json:"timestamp"`

	// **参数解释**： 主机配置信息 **取值范围**： 不涉及
	Value string `json:"value"`
}

func (o EventDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDataInfo struct{}"
	}

	return strings.Join([]string{"EventDataInfo", string(data)}, " ")
}
