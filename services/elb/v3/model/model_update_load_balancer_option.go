package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新负载均衡器参数。
type UpdateLoadBalancerOption struct {
	// 负载均衡器的名称。

	Name *string `json:"name,omitempty"`
	// 负载均衡器的管理状态。只能设置为true。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 负载均衡器的描述。

	Description *string `json:"description,omitempty"`
	// 双栈类型负载均衡器所在子网的IPv6网络ID。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。 通过更新ipv6_vip_virsubnet_id可以更新负载均衡器所在IPv6子网，并且负载均衡器的内网IPv6地址将发生变化。 ipv6_vip_virsubnet_id对应的子网必须属于当前负载均衡器所在VPC。 注： 1.只有子网开启IPv6时才可以传入。 2.仅当guaranteed是true的场合，才支持更新。 3.传入为null表示解绑IPv6子网。 [不支持IPv6，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器所在的IPv4子网ID。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。 通过更新vip_subnet_cidr_id可以更新负载均衡器所在IPv4子网，并且负载均衡器的内网IPv4地址将发生变化。 若同时设置了vip_address，则必须保证vip_address对应的IP在vip_subnet_cidr_id的子网网段中。更新后将使用vip_address对应的IP作为负载均衡器的内网IPv4地址。 vip_subnet_cidr_id对应的子网必须属于当前负载均衡器vpc_id对应的VPC。 注： 1.只有guaranteed是true的负载均衡器才支持更新vip_subnet_cidr_id。 2.传入null表示解绑IPv4子网。

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`
	// 负载均衡器的IPv4虚拟IP。该地址必须包含在所在子网的IPv4网段内，且未被占用。  注：仅当guaranteed是true的场合，才支持更新。

	VipAddress *string `json:"vip_address,omitempty"`
	// 四层Flavor ID。  注： 1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null，只允许改大，不允许改小。  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso)

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`
	// 七层Flavor ID。  注： 1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null；只允许改大，不允许改小。  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso)

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// 是否启用跨VPC后端转发，值只允许为true。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 下联面子网的网络ID列表。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。 若指定多个下联面子网，则按顺序优先使用第一个子网来为负载均衡器下联面端口分配ip地址。 该参数将全量更新LB的下联面子网，即若LB原有的下联面子网网络ID不在该字段数组中，则将移除LB与该下联面子网的关联关系。但不允许移除已被ELB使用的子网。 使用说明： - 所有ID同属于该LB所在的VPC。 - 不支持边缘云子网。

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
	// 是否开启删除保护。取值：false不开启，true开启。 > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用](tag:otc,otc_test,dt,dt_test)

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
