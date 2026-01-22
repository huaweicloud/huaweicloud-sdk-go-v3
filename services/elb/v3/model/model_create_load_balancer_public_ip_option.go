package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLoadBalancerPublicIpOption struct {

	// **参数解释**：IP版本。  **约束限制**：不涉及  **取值范围**：  - 4：表示IPv4。  - 6：表示IPv6。  **默认取值**：4  [不支持IPv6，请勿设置为6。](tag:dt)
	IpVersion *int32 `json:"ip_version,omitempty"`

	// **参数解释**：弹性公网IP的网络类型，默认5_bgp，更多请参考弹性公网ip创建。  [**约束限制**：华南-深圳局点该参数取值只能为5_gray。](tag:hws) [**约束限制**：只支持设置为5_gray。](tag:dt)  **取值范围**：5_bgp - 全动态BGP。  **默认取值**：5_bgp
	NetworkType string `json:"network_type"`

	// **参数解释**：计费订单信息。  **约束限制**： - 空：按需计费。 [- 非空：包周期计费。格式：order_id:product_id:region_id:project_id](tag:hws)  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,srg,fcs,fcs_vm,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,ct)
	BillingInfo *string `json:"billing_info,omitempty"`

	// **参数解释**：弹性公网IP的描述信息。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	Bandwidth *CreateLoadBalancerBandwidthOption `json:"bandwidth"`
}

func (o CreateLoadBalancerPublicIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerPublicIpOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerPublicIpOption", string(data)}, " ")
}
