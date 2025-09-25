package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAutoEnlargePolicyResponse Response Object
type SearchAutoEnlargePolicyResponse struct {

	// **参数解释**: 磁盘自动扩容开关。 **取值范围**: 不涉及。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// **参数解释**: 存储自动扩容上限。 **取值范围**: 不涉及。
	LimitVolumeSize *int32 `json:"limit_volume_size,omitempty"`

	// **参数解释**: 最小扩容磁盘容量。 **取值范围**: 不涉及。
	MinVolumeSize *int32 `json:"min_volume_size,omitempty"`

	// **参数解释**: 最大扩容磁盘容量。 **取值范围**: 不涉及。
	MaxVolumeSize *int32 `json:"max_volume_size,omitempty"`

	// **参数解释**: 可用存储空间率。 **取值范围**: 不涉及。
	TriggerAvailablePercent *int32 `json:"trigger_available_percent,omitempty"`

	// **参数解释**: 空间率集合。 **取值范围**: 不涉及。
	Percents *[]int32 `json:"percents,omitempty"`

	// **参数解释**: 扩容步长，固定大小扩容方式。 **取值范围**: 不涉及。
	StepSize *int32 `json:"step_size,omitempty"`

	// **参数解释**: 扩容步长，百分比扩容方式。 **取值范围**: 不涉及。
	StepPercent    *int32 `json:"step_percent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"SearchAutoEnlargePolicyResponse", string(data)}, " ")
}
