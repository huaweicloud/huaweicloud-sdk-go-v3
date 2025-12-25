package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdditionalInfo **参数解释**： 告警记录额外字段，仅针对事件监控告警场景所产生的告警记录信息。
type AdditionalInfo struct {

	// **参数解释**： 该条告警记录对应的资源ID；如：22d98f6c-16d2-4c2d-b424-50e79d82838f。 **取值范围**： 长度为[0,128]个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 该条告警记录对应的资源名称；如：ECS-Test01。 **取值范围**： 长度为[0,128]个字符。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 该条告警记录对应的事件监控ID，资源所产生的事件；如：ev16031292300990kKN8p17J。 **取值范围**： 字符串长度为24。
	EventId *string `json:"event_id,omitempty"`
}

func (o AdditionalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalInfo struct{}"
	}

	return strings.Join([]string{"AdditionalInfo", string(data)}, " ")
}
