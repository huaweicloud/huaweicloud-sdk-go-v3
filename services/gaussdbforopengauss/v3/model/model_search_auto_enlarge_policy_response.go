package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAutoEnlargePolicyResponse Response Object
type SearchAutoEnlargePolicyResponse struct {

	// 磁盘自动扩容开关。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// 存储自动扩容上限。
	LimitVolumeSize *int32 `json:"limit_volume_size,omitempty"`

	// 最小扩容磁盘容量。
	MinVolumeSize *int32 `json:"min_volume_size,omitempty"`

	// 最大扩容磁盘容量。
	MaxVolumeSize *int32 `json:"max_volume_size,omitempty"`

	// 可用存储空间率。
	TriggerAvailablePercent *int32 `json:"trigger_available_percent,omitempty"`

	// 空间率集合。
	Percents *[]int32 `json:"percents,omitempty"`

	// 扩容步长，固定大小扩容方式。
	StepSize *int32 `json:"step_size,omitempty"`

	// 扩容步长，百分比扩容方式。
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
