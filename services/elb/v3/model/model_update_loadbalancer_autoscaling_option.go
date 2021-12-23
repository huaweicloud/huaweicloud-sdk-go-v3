package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性扩缩容配置信息。负载均衡器配置并开启弹性扩缩容后，可根据负载情况自动调整负载均衡器的规格。  使用说明： - 仅当局点支持弹性扩缩特性该字段才有效。可咨询客服了解当前局点是否支持弹性扩缩容特性。 - 开启弹性扩缩容后，l4_flavor_id和l7_flavor_id不再起作用。
type UpdateLoadbalancerAutoscalingOption struct {
	// 当前负载均衡器是否开启弹性扩缩容。 取值： - true：开启。 - false：不开启。

	Enable bool `json:"enable"`
	// 弹性扩缩容的最小四层规格ID，有四层监听器时，该字段不能为空。

	MinL4FlavorId *string `json:"min_l4_flavor_id,omitempty"`
	// 弹性扩缩容的最小七层规格ID，有七层监听器时，该字段不能为空。

	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o UpdateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerAutoscalingOption", string(data)}, " ")
}
