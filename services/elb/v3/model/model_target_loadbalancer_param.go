package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetLoadbalancerParam 负载均衡器复制新实例相关配置参数。
type TargetLoadbalancerParam struct {

	// 新实例名称。 可选，不选时使用源负载均衡器名称加copy-x的后缀作为新实例名称。
	Name *string `json:"name,omitempty"`

	// 新实例所属可用区。 可选，不选时使用源负载均衡器的可用区。 只在独享型复制场景可配置。
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	// 新实例所属子网的ipv4子网id。 可选，不选时使用源负载均衡器的ipv4子网。 所选子网需要与源负载均衡器在同一个vpc内。
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// 新实例的ipv4私网地址。 可选，不选时随机分配。 只在独享型复制场景、共享型复制为独享型场景可配。
	VipAddress *string `json:"vip_address,omitempty"`

	// 新实例ipv6网络所属的子网网络id。 可选，不选时使用源负载均衡器的子网。 所选子网需要与源负载均衡器在同一个vpc内。 只在独享型复制场景可配。
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// 新实例的ipv6地址。 可选，不选时随机分配。 只在独享型复制场景可配。
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// 新实例后端子网的网络id。 可选，不选时使用源负载均衡器的后端子网。 所选子网需要与源负载均衡器在同一个vpc内。 只在独享型复制场景、共享型复制为独享型场景可配。
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// 新实例4层规格。 可选，不选时使用源负载均衡器的4层规格。 只在独享型复制场景、共享型复制为独享型场景可配。
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// 新实例7层规格。 可选，不选时使用源负载均衡器的7层规格。 只在独享型复制场景、共享型复制为独享型场景可配。
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// 新实例所属企业项目。 可选，不选时使用源负载均衡器的企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 新实例是否复用源ELB的后端服务器组和后端服务器标识。 可选，配置为true时需要开启后端服务器组多实例挂载功能。 不选时默认新创建后端服务器组。 enterprise_project_id选项配置使用其他企业项目时，该选项失效。 只在独享型复制场景、共享型复制为独享型场景可配。
	ReusePool *bool `json:"reuse_pool,omitempty"`

	// 新实例类型。 可选配置。 独享型复制场景默认为true，若显式指定，只能配置为true。 共享型复制场景默认为false，若显式指定，配置为false表示新复制共享型实例，配置为true表示新复制独享型实例。
	Guaranteed *bool `json:"guaranteed,omitempty"`
}

func (o TargetLoadbalancerParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetLoadbalancerParam struct{}"
	}

	return strings.Join([]string{"TargetLoadbalancerParam", string(data)}, " ")
}
