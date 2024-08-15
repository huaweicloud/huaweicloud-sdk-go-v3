package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadBalancerPublicIpOption 参数解释：创建ELB时，新建公网IP请求参数
type CreateLoadBalancerPublicIpOption struct {

	// 参数解释：IP版本。  约束限制：取值只有4和6。4：IPv4, 6: IPv6。 [不支持IPv6，请勿设置为6。](tag:dt,dt_test)  取值范围：  - 4：表示IPv4。  - 6：表示IPv6。  默认取值：4
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 参数解释：弹性公网IP的网络类型，默认5_bgp，更多请参考弹性公网ip创建。  [约束限制：华南-深圳局点该参数取值只能为5_gray。](tag:hws) [约束限制：只支持设置为5_gray。](tag:dt)  取值范围：5_bgp。  默认取值：5_bgp。
	NetworkType string `json:"network_type"`

	// 参数解释：计费订单信息。  约束限制： - 空：按需计费。 [ - 非空：包周期计费。格式：order_id:product_id:region_id:project_id ](tag:hws,hk,hws_eu,otc,tlf,ctc,hcso,sbc,g42,cmcc,hk_g42,dt_test,hcso_dt,mix,hk_sbc,hws_ocb,hk_vdf,fcs,fcs_dt,dt)  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	BillingInfo *string `json:"billing_info,omitempty"`

	// 参数解释：弹性公网IP的描述信息。
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
