package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TriggerMetadataList struct {

	// 触发名称
	TriggerName *string `json:"trigger_name,omitempty"`

	// 触发器类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 事件数据
	EventData *string `json:"event_data,omitempty"`
}

func (o TriggerMetadataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerMetadataList struct{}"
	}

	return strings.Join([]string{"TriggerMetadataList", string(data)}, " ")
}
