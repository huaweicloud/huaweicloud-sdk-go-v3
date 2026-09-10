package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachDevServerPortsRequestBody struct {

	// **参数解释**：网卡ID，填该参数时，表明挂载已有网卡，其他参数不用填。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PortId *string `json:"port_id,omitempty"`

	// **参数解释**：网卡名称。 **约束限制**：不涉及。 **取值范围**：默认为空，最大长度不超过255。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：端口子网ID。 **约束限制**：参数port_id未填时，需要新建网卡进行挂载，此时network_id为必填项。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	NetworkId *string `json:"network_id,omitempty"`

	// **参数解释**：端口IP地址。 **约束限制**：不支持更新。 **取值范围**：所属网络网段。 **默认取值**：不涉及。
	IpAddress *string `json:"ip_address,omitempty"`

	// **参数解释**：关联安全组ID列表。 **约束限制**：一个端口默认最多吃吃100个安全组。 **默认取值**：不涉及。
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// **参数解释**：是否使能efi。 **约束限制**：不涉及。 **取值范围**： - true：启用efi - false：不启用efi  **默认取值**：不涉及。
	EnableEfi *bool `json:"enable_efi,omitempty"`

	// **参数解释**：efi 协议。 **约束限制**：不涉及。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	EfiProtocol *string `json:"efi_protocol,omitempty"`
}

func (o AttachDevServerPortsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDevServerPortsRequestBody struct{}"
	}

	return strings.Join([]string{"AttachDevServerPortsRequestBody", string(data)}, " ")
}
