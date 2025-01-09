package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdditionalInfo 告警记录额外字段，仅针对事件监控告警场景所产生的告警记录信息。
type AdditionalInfo struct {

	// 该条告警记录对应的资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 该条告警记录对应的资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 该条告警记录对应的事件监控ID，资源所产生的事件
	EventId *string `json:"event_id,omitempty"`
}

func (o AdditionalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalInfo struct{}"
	}

	return strings.Join([]string{"AdditionalInfo", string(data)}, " ")
}
