package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberOption **参数解释**：更新后端服务器请求参数。  **约束限制**：不涉及
type UpdateMemberOption struct {

	// **参数解释**：后端服务器的管理状态。 虽然创建、更新请求支持该字段，但实际取值决定于后端服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。  **约束限制**：请勿传入该字段。  **取值范围**：true、false  **默认取值**：不涉及
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：后端服务器的可用区。  **约束限制**： 仅支持IP类型后端服务器更新该字段。且后端服务器组开启可用区亲和时，IP类型后端服务器必须配置该字段，且无法更新为\"\"。  **取值范围**：本region中ECS可选择的可用区。  **默认取值**：不涉及
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：后端服务器名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：后端服务器的权重，请求将根据pool配置的负载均衡算法和后端服务器的权重进行负载分发。权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  **约束限制**：若所在pool的lb_algorithm取值为SOURCE_IP或QUIC_CID，该字段无效。  **取值范围**：0-100  **默认取值**：1
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释**：后端服务器端口。  **约束限制**： - 在开启端口透传的pool下的member，该字段无法更新。 - 网关型LB，即pool协议为IP时，protocol_port必须设置为0。  **默认取值**：不涉及
	ProtocolPort *int32 `json:"protocol_port,omitempty"`
}

func (o UpdateMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberOption struct{}"
	}

	return strings.Join([]string{"UpdateMemberOption", string(data)}, " ")
}
