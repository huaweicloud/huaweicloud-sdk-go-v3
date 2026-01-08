package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountPreoccupyIpNumRequest Request Object
type CountPreoccupyIpNumRequest struct {

	// **参数解释**：负载均衡器七层规格的ID。传入该字段表示计算创建该规格的LB的预占IP数量，或变更LB的原七层规格到该规格所需要的新增预占IP数量。  **约束限制**：仅支持创建LB、变更LB规格场景。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持传入l7_flavor_id。](tag:hcso,hk_vdf,srg,fcs)
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// **参数解释**：IP类型后端转发开关。  **约束限制**：仅支持创建LB、LB开启IP类型后端转发场景。  **取值范围**： - 取值true表示计算创建开启IP类型后端转发的LB的预占IP数量，或开启IP类型后端转发所需要的新增预占IP数量。 - 取值false表示计算创建不开启IP类型后端转发的LB的预占IP。 - 不传等价false。  **默认取值**：false  [荷兰region不支持该字段，请勿使用。](tag:dt)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// **参数解释**：负载均衡器IP地址类型，取值4、6 。  **约束限制**：仅支持创建LB场景。  **取值范围**： - 取值4表示计算创建支持IPv4地址的LB的预占IP。 - 取值6表示计算创建支持IPv6地址的LB的预占IP。  **默认取值**：不涉及  [不支持IPv6，请勿设置为6。](tag:dt)
	IpVersion *int32 `json:"ip_version,omitempty"`

	// **参数解释**：负载均衡器ID。计算LB变更或创建LB中的第一个七层监听器的新增预占IP。  **约束限制**：仅支持变更LB规格、开启IP类型后端转发、开启/不开启地址转换功能、创建LB中的第一个七层监听器场景。  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// **参数解释**：计算创建AZ列表为availability_zone_id的负载均衡器实例的预占IP。  **约束限制**： - 仅支持创建LB场景。 - 传入loadbalancer_id时，该参数无效。  **取值范围**：不涉及  **默认取值**：不涉及
	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`

	// **参数解释**：计算共享型升级为独享型ELB负载均衡器实例的所需预占IP。  **约束限制**： - 仅支持共享型升级为独享型ELB场景。 - 必须同时传入loadbalancer_id。  **取值范围**：UPGRADE - 共享型升级为独享型ELB场景。  **默认取值**：不涉及
	Scene *string `json:"scene,omitempty"`

	// **参数解释**：开启地址转换。传入该字段表示计算创建LB及该LB下开启/不开启地址转换特性的监听器所需要的预占IP，或者指定LB创建开启/不开启地址转换特性的监听器所需要的新增预占IP。  **约束限制**： - 仅支持计算ELB实例开启NAT64特性场景。  **取值范围**： - true：开启地址转换特性。 - false：不开启地址转换特性。  **默认取值**：false
	Nat64Enable *bool `json:"nat64_enable,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumRequest struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}
