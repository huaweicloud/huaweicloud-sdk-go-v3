package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetLoadbalancerParam **参数解释**：负载均衡器复制新实例参数对象。  **约束限制**：不涉及
type TargetLoadbalancerParam struct {

	// **参数解释**：新实例名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不传时使用源ELB名称加\"-copy-{x}\"的后缀作为新实例名称。{x}代表数字序号。
	Name *string `json:"name,omitempty"`

	// **参数解释**：新实例所属可用区。  **约束限制**：仅支持源ELB独享型复制场景设置该字段。  **取值范围**：不涉及  **默认取值**：不传时使用源ELB的availability_zone_list。
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	// **参数解释**：新实例所属子网的ipv4子网ID。  **约束限制**：所选子网必须与源ELB在同一个vpc内。  **取值范围**：不涉及  **默认取值**：不传时使用源ELB的vip_subnet_cidr_id。
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// **参数解释**：新实例的ipv4私网地址。  **约束限制**：仅支持源ELB独享型复制场景、源ELB共享型复制为独享型场景支持设置该字段。  **取值范围**：不涉及  **默认取值**：不传时随机分配vip_subnet_cidr_id对应子网的可用IP地址。
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**：新实例ipv6网络所属的子网网络ID。  **约束限制**： - 仅支持源ELB独享型复制场景设置该字段。 - 所选子网必须与源负载均衡器在同一个vpc内。  **取值范围**：不涉及  **默认取值**：不传时使用源ELB的ipv6_vip_virsubnet_id。
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// **参数解释**：新实例的ipv6地址。  **约束限制**：仅支持源ELB为独享型复制场景设置该字段。  **取值范围**：不涉及  **默认取值**：不传时随机分配ipv6_vip_virsubnet_id对应子网的可用IP地址。
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// **参数解释**：新实例后端子网的网络ID。  **约束限制**： - 仅支持源ELB独享型复制场景、源ELB共享型复制为独享型场景支持设置该字段。 - 所选子网必须与源负载均衡器在同一个vpc内。  **取值范围**：不涉及  **默认取值**：不传时使用源负载均衡器的后端子网。
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// **参数解释**：新实例4层规格。  **约束限制**：仅支持源ELB独享型复制场景、源ELB共享型复制为独享型场景支持设置该字段。  **取值范围**：不涉及  **默认取值**：不传时保持与源负载均衡器的4层规格一致。
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// **参数解释**：新实例7层规格。  **约束限制**：仅支持源ELB独享型复制场景、源ELB共享型复制为独享型场景支持设置该字段。  **取值范围**：不涉及  **默认取值**：不传时保持与源负载均衡器的7层规格一致。
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// **参数解释**：资源所属的企业项目ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不传时保持与源负载均衡器的企业项目一致。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：新实例是否复用源ELB的后端服务器组。  **约束限制**： - 设置为true时，需要开启后端服务器组的多实例挂载功能。 - 请求参数enterprise_project_id使用与源ELB不同的其他企业项目时，该参数失效。 - 仅源ELB独享型复制场景、源ELB共享型复制为独享型场景支持设置为true。  **取值范围**： - false：新创建后端服务器组。 - true：复用源ELB的后端服务器组。  **默认取值**：false
	ReusePool *bool `json:"reuse_pool,omitempty"`

	// **参数解释**：新实例类型。  **约束限制**：不涉及  **取值范围**： - false：复制为共享型实例，此时源ELB必须是共享型。 - true：复制为独享型实例，源ELB可以是共享型或独享型。  **默认取值**： - 源ELB是独享型复制场景默认为true。 - 源ELB是共享型复制场景默认为false。
	Guaranteed *bool `json:"guaranteed,omitempty"`
}

func (o TargetLoadbalancerParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetLoadbalancerParam struct{}"
	}

	return strings.Join([]string{"TargetLoadbalancerParam", string(data)}, " ")
}
