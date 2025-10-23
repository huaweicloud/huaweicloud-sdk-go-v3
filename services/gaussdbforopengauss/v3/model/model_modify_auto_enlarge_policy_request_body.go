package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyAutoEnlargePolicyRequestBody struct {

	// **参数解释**: 磁盘自动扩容开关。 **约束限制**: 不涉及。 **取值范围**: - true：开启磁盘自动扩容。 - false：关闭磁盘自动扩容。  **默认取值**: 不涉及。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// **参数解释**: 存储自动扩容上限。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	LimitVolumeSize *int32 `json:"limit_volume_size,omitempty"`

	// **参数解释**: 可用存储空间率。 **约束限制**: 不涉及。 **取值范围**: 0-100。 **默认取值**: 20。
	TriggerAvailablePercent *int32 `json:"trigger_available_percent,omitempty"`

	// **参数解释**: 扩容步长，固定大小扩容方式。 **约束限制**: 40的倍数。 **取值范围**: 不涉及。 **默认取值**: 40。
	StepSize *int32 `json:"step_size,omitempty"`

	// **参数解释**: 扩容步长，百分比扩容方式。 **约束限制**: 不能小于1。 **取值范围**: 不涉及。 **默认取值**: 20。
	StepPercent *int32 `json:"step_percent,omitempty"`
}

func (o ModifyAutoEnlargePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoEnlargePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyAutoEnlargePolicyRequestBody", string(data)}, " ")
}
