package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingRef 弹性扩缩容配置信息。负载均衡器配置并开启弹性扩缩容后，可根据负载情况自动调整负载均衡器的规格。  使用说明： - 仅当租户白名单放开后该字段才有效 - 开启弹性扩缩容后，l4_flavor_id和l7_flavor_id表示该LB实例弹性规格的上限。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,hcso,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
type AutoscalingRef struct {

	// 当前负载均衡器是否开启弹性扩缩容。  取值： - true：开启。 - false：不开启，默认值。
	Enable bool `json:"enable"`

	// 弹性扩缩容的最小七层规格ID（规格类型L7_elastic），有七层监听器时，该字段不能为空。
	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o AutoscalingRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingRef struct{}"
	}

	return strings.Join([]string{"AutoscalingRef", string(data)}, " ")
}
