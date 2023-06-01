package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建ELB时，新建公网IP请求参数
type CreateLoadBalancerPublicIpOption struct {

	// IP版本。  取值：4表示IPv4，6表示IPv6。  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 弹性公网IP的网络类型，默认5_bgp，更多请参考弹性公网ip创建。  [华南-深圳局点该参数取值只能为5_gray](tag:hws) [只支持设置为5_gray](tag:dt)
	NetworkType string `json:"network_type"`

	// 资源账单信息。  取值： - 空：按需计费。 - 非空：包周期计费。  [不支持该字段，请勿使用](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,hcso,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	BillingInfo *string `json:"billing_info,omitempty"`

	// 弹性公网IP的描述信息，不支持特殊字符
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
