package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性扩缩容配置信息。负载均衡器配置并开启弹性扩缩容后，可根据负载情况自动调整负载均衡器的规格。   使用说明： - 仅当租户白名单放开后该字段才有效 - 开启弹性扩缩容后，l4_flavor_id和l7_flavor_id表示该LB实例弹性规格的上限。
type CreateLoadbalancerAutoscalingOption struct {

	// 负载均衡器弹性扩缩容开关
	Enable bool `json:"enable" xml:"enable"`

	// 弹性扩缩容的最小七层规格ID，类型为L7_Elastic，有七层监听器时，该字段不能为空。
	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty" xml:"min_l7_flavor_id"`
}

func (o CreateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerAutoscalingOption", string(data)}, " ")
}
