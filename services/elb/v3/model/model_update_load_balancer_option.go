package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新负载均衡器参数。
type UpdateLoadBalancerOption struct {
	// 负载均衡器的名称。

	Name *string `json:"name,omitempty"`
	// 负载均衡器的管理状态。只能设置为true。  [不支持该字段，请勿使用。](tag:dt,dt_test)

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 负载均衡器的描述。

	Description *string `json:"description,omitempty"`
	// 双栈类型负载均衡器所在子网的IPv6网络ID。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。  通过更新ipv6_vip_virsubnet_id可以更新负载均衡器所在IPv6子网，且负载均衡器的内网IPv6地址将发生变化。  使用说明： - ipv6_vip_virsubnet_id 对应的子网必须属于当前负载均衡器所在VPC，且子网需开启支持IPv6。 - 只有guaranteed是true的负载均衡器才支持更新ipv6_vip_virsubnet_id。 - *传入为null表示解绑IPv6子网。* - 更新IPv6子网不会导致IPv4子网发生变化。  [不支持IPv6，请勿使用。](tag:dt,dt_test)

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器所在的IPv4子网ID。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。 通过更新vip_subnet_cidr_id可以更新负载均衡器所在IPv4子网，并且负载均衡器的内网IPv4地址将发生变化。 使用说明： - 若同时设置了vip_address，则必须保证vip_address对应的IP在vip_subnet_cidr_id的子网网段中。 - vip_subnet_cidr_id对应的子网必须属于当前负载均衡器vpc_id对应的VPC。 - 只有guaranteed是true的负载均衡器才支持更新vip_subnet_cidr_id。 - *传入null表示解绑IPv4子网。* - 更新IPv子网不会导致IPv4子网发生变化。

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`
	// 负载均衡器的IPv4虚拟IP。该地址必须包含在所在子网的IPv4网段内，且未被占用。  注：仅当guaranteed是true的场合，才支持更新。

	VipAddress *string `json:"vip_address,omitempty"`
	// 四层Flavor ID。  注： 1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null，只允许改大，不允许改小。  [hcso场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hws,hcso)

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`
	// 七层Flavor ID。  注： 1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null；只允许改大，不允许改小。  [hcso场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hws,hcso)

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// 是否启用跨VPC后端转发，开启跨VPC后端转发后，支持添加其他VPC、其他公有云、云下数据中心的服务器。取值： - true：开启。 - false：不开启。 [仅独享型负载均衡器支持该特性，且只能更新为true，即开启后不支持关闭。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test)  [不支持该字段，请勿使用。](tag:dt,dt_test)

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 下联面子网的网络ID列表。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。  已绑定的下联面子网也在传参elb_virsubnet_ids列表中，则绑定关系保留。  已绑定的下联面子网若不在传参elb_virsubnet_ids列表中，则将移除LB与该下联面子网的关联关系。但不允许移除已被ELB使用的子网，否则将报错，不做任何修改。  在传参elb_virsubnet_ids列表中但不在已绑定的下联面子网列表中，则将新增LB与下联面的绑定关系。   使用说明：  - 所有ID同属于该LB所在的VPC。  - 不支持边缘云子网。

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
	// 是否开启删除保护。取值：false不开启，true开启。 > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用](tag:dt,dt_test)

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidUpdateOption `json:"prepaid_options,omitempty"`

	Autoscaling *UpdateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`
}

func (o UpdateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerOption", string(data)}, " ")
}
