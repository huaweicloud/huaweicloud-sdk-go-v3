package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingRef **参数解释**：弹性扩缩容配置信息。负载均衡器开启弹性扩缩容后，可根据负载情况自动调整负载均衡器的规格。  [> 该字段已经废弃，但仍然保留兼容性支持，建议不要使用该字段。如果传入该字段，创建的弹性实例将会有保底规格并产生对应保底规格的费用。](tag:cmcc,ctc,dc2,dt,fcs_arm,fcs_dt,fm,h3dc,hcs,hcso_dt,HEC,hk_sbc,hk_tm,hws,hws_hk,hws_ocb,hws_test,kvm,nohcs,ocb,ct,sbc,sfsturb,tlf,tlf_test,tm)  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,hcso,hk_vdf,srg,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
type AutoscalingRef struct {

	// **参数解释**：当前负载均衡器是否开启弹性扩缩容。  **取值范围**： - true：开启。 - false：不开启。
	Enable bool `json:"enable"`

	// **参数解释**：弹性扩缩容的最小七层规格ID（规格类型L7_elastic）。  **取值范围**：不涉及
	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o AutoscalingRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingRef struct{}"
	}

	return strings.Join([]string{"AutoscalingRef", string(data)}, " ")
}
