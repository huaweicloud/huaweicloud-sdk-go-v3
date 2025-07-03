package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateSpotOptions 竞价计费信息
type TemplateSpotOptions struct {

	// 用户愿意为竞价实例每小时支付的最高价格
	SpotPrice *float32 `json:"spot_price,omitempty"`

	// 购买的竞价实例时长
	BlockDurationMinutes *int32 `json:"block_duration_minutes,omitempty"`

	// 竞价实例中断策略，当前支持immediate
	InstanceInterruptionBehavior *string `json:"instance_interruption_behavior,omitempty"`
}

func (o TemplateSpotOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateSpotOptions struct{}"
	}

	return strings.Join([]string{"TemplateSpotOptions", string(data)}, " ")
}
