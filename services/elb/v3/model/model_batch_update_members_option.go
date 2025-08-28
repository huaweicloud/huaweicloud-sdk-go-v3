package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateMembersOption struct {

	// **参数解释**：后端服务器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  >此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id"`

	// **参数解释**：后端服务器的可用区。  **约束限制**：仅支持IP类型后端服务器设置该字段。且后端服务器组开启可用区亲和时，IP类型后端服务器必须配置该字段为有效非空值。  **取值范围**：本region中ECS可选择的可用区。  **默认取值**：不涉及
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：后端服务器的管理状态。取值：true、false。  **约束限制**：虽然创建、更新请求支持该字段，但实际取值决定于后端服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。  **取值范围**：不涉及  **默认取值**：不涉及  请勿传入该字段。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：后端服务器名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：后端服务器端口。  **约束限制**： - 在开启端口透传的pool下的member，该字段无法更新。 [- 网关型LB，即pool协议为IP时，protocol_port必须设置为0。](tag:hws_eu)  **取值范围**：1-65535  **默认取值**：不涉及
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// **参数解释**：后端服务器的权重，请求按权重在同一后端服务器组下的后端服务器间分发。权重为0的后端不再接受新的请求。  **约束限制**：当后端服务器所在的后端服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。  **取值范围**：0-100  **默认取值**：不涉及
	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchUpdateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersOption", string(data)}, " ")
}
