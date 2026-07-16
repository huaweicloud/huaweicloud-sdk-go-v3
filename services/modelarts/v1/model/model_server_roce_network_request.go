package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerRoceNetworkRequest struct {

	// **参数解释**：RoCE网络类型。 **约束限制**：不涉及。 **取值范围**：  - vxlan_roce  - roce_v2  **默认取值**：不涉及。
	NetworkType *string `json:"network_type,omitempty"`

	// **参数解释**：物理网络名称。 **约束限制**：^[-_.a-zA-Z0-9]{1,64}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PhysicalNetwork *string `json:"physical_network,omitempty"`
}

func (o ServerRoceNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerRoceNetworkRequest struct{}"
	}

	return strings.Join([]string{"ServerRoceNetworkRequest", string(data)}, " ")
}
