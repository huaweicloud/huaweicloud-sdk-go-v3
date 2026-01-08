package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerNode
type LoadBalancerNode struct {

	// **参数解释**：负载均衡器id。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：负载均衡器名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：是否独享型LB。  **取值范围**： - false：共享型。 - true：独享型。
	Guaranteed bool `json:"guaranteed"`

	// **参数解释**：网络型规格ID。  **取值范围**：不涉及
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// **参数解释**：应用型规格ID。  **取值范围**：不涉及
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// **参数解释**：ipv4地址。  **取值范围**：不涉及
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**：ipv6地址。  **取值范围**：不涉及
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// **参数解释**：负载均衡器所在的可用区列表。  **取值范围**：不涉及
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`
}

func (o LoadBalancerNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerNode struct{}"
	}

	return strings.Join([]string{"LoadBalancerNode", string(data)}, " ")
}
