package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerRoceNetwork struct {

	// **参数解释**：RoCE网络id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：RoCE网络名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：项目ID。 **取值范围**：不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：状态。 **参数解释**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：租户id。 **取值范围**：不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**：子网。 **取值范围**：不涉及。
	Subnets *string `json:"subnets,omitempty"`

	// **参数解释**：RoCE网络类型。 **取值范围**：不涉及。
	ProvidernetworkType *string `json:"provider:network_type,omitempty"`

	// **参数解释**：实际物理网络。 **取值范围**：不涉及。
	ProviderphysicalNetwork *string `json:"provider:physical_network,omitempty"`
}

func (o ServerRoceNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerRoceNetwork struct{}"
	}

	return strings.Join([]string{"ServerRoceNetwork", string(data)}, " ")
}
