package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建负载均衡器参数。
type CreateLoadBalancerOption struct {
	// 负载均衡器的名称。

	Name *string `json:"name,omitempty"`
	// 负载均衡器的描述。

	Description *string `json:"description,omitempty"`
	// 负载均衡器的IPv4虚拟IP。该地址必须包含在所在子网的IPv4网段内，且未被占用。  使用说明： - 传入vip_address时必须传入vip_subnet_cidr_id。 - 不传入vip_address，但传入vip_subnet_cidr_id，则自动分配IPv4虚拟IP。 - 不传入vip_address，且不传vip_subnet_cidr_id，则不分配虚拟IP，vip_address=null。

	VipAddress *string `json:"vip_address,omitempty"`
	// 负载均衡器所在子网的IPv4子网ID。若需要创建带IPv4虚拟IP的LB，该字段必须传入。 可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。 使用说明： - vpc_id, vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 若同时传入vpc_id和vip_subnet_cidr_id，则vip_subnet_cidr_id对应的子网必须属于vpc_id对应的VPC。

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`
	// 双栈类型负载均衡器所在子网的IPv6网络ID。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。  使用说明： - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 需要对应的子网开启IPv6。  [不支持IPv6，请勿使用](tag:otc,otc_test,dt,dt_test)

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器的生产者名称。固定为vlb。

	Provider *string `json:"provider,omitempty"`
	// 四层Flavor ID。[创建负载均衡器时l4_flavor_id和l7_flavor_id不能都不传](tag:otc,otc_test,dt,dt_test) [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso) 注意：当l4_flavor_id和l7_flavor_id都不传的时，会选择默认flavor。

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`
	// 负载均衡器所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 是否独享型负载均衡器。取值： - true：独享型。 - false：共享型。  当前只支持设置为true，设置为false会返回400 Bad Request 。默认：true。

	Guaranteed *bool `json:"guaranteed,omitempty"`
	// 负载均衡器所在的VPC ID。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/vpcs 响应参数中的id得到。  使用说明： - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。

	VpcId *string `json:"vpc_id,omitempty"`
	// 可用区列表。可通过GET https&#58;//{ELB_Endponit}/v3/{project_id}/elb/availability-zones接口来查询可用区集合列表。创建负载均衡器时，从查询结果选择某一个可用区集合，并从中选择一个或多个可用区。

	AvailabilityZoneList []string `json:"availability_zone_list"`
	// 负载均衡器所属的企业项目ID。不能传入\"\"、\"0\"或不存在的企业项目ID，创建时不传则资源属于default企业项目，默认返回\"0\"。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 负载均衡的标签列表。示例：\"tags\":[{\"key\":\"my_tag\",\"value\":\"my_tag_value\"}]

	Tags *[]Tag `json:"tags,omitempty"`
	// 负载均衡器的管理状态。只能设置为true。默认：true。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 七层Flavor ID。[创建负载均衡器时l4_flavor_id和l7_flavor_id不能都不传](tag:otc,otc_test,dt,dt_test) [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso) 注意：当l4_flavor_id和l7_flavor_id都不传的时，会选择默认flavor。

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`
	// 资源账单信息，取值： - 空：按需计费。 - 非空：包周期计费。 包周期计费billing_info字段的格式为：order_id&#58;product_id&#58;region_id&#58;project_id，如： CS2107161019CDJZZ&#58;OFFI569702121789763584&#58;eu-de&#58;057ef081eb00d2732fd1c01a9be75e6f 使用说明： - admin权限才能更新此字段。 [不支持该字段，请勿使用](tag:otc,otc_test,dt,dt_test)

	BillingInfo *string `json:"billing_info,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// 负载均衡器绑定的公网IP ID。只支持绑定数组中的第一个EIP，其他将被忽略。

	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	Publicip *CreateLoadBalancerPublicIpOption `json:"publicip,omitempty"`
	// 下联面子网的网络ID列表。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets  响应参数中的id得到。 若不指定该字段，则会在当前负载均衡器所在的VPC中任意选一个子网，优选双栈网络。 若指定多个下联面子网，则按顺序优先使用第一个子网来为负载均衡器下联面端口分配ip地址。 下联面子网必须属于该LB所在的VPC。

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
	// 是否启用跨VPC后端转发。取值：true 表示开启，false 表示不开启。默认：false不开启。仅独享型负载均衡器支持该特性。  开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 是否开启删除保护。取值：false不开启，true开启。默认false不开启。 > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用](tag:otc,otc_test,dt,dt_test)

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidCreateOption `json:"prepaid_options,omitempty"`

	Autoscaling *CreateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`
}

func (o CreateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerOption", string(data)}, " ")
}
