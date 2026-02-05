package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdmGroupInfo struct {

	// **参数解释**：  组ID。  **参数范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  名称。  **参数范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  角色。  **参数范围**：  不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释**：  组ip。  **参数范围**：  不涉及。
	Endpoint *string `json:"endpoint,omitempty"`

	// **参数解释**：  ipv6。  **参数范围**：  不涉及。
	Ipv6Endpoint *string `json:"ipv6_endpoint,omitempty"`

	// **参数解释**：  节点数量。  **参数范围**：  不涉及。
	NodeCount *int32 `json:"node_count,omitempty"`

	// **参数解释**：  负载均衡。  **参数范围**：  不涉及。
	LoadBalance *bool `json:"load_balance,omitempty"`

	// **参数解释**：  是否默认组。  **参数范围**：  不涉及。
	IsDefaultGroup *bool `json:"is_default_group,omitempty"`

	// **参数解释**：  其他信息。  **参数范围**：  不涉及。
	ExtendMap map[string]string `json:"extend_map,omitempty"`

	// **参数解释**：  节点信息。  **参数范围**：  不涉及。
	Nodes *[]DdmNodeInfo `json:"nodes,omitempty"`
}

func (o DdmGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdmGroupInfo struct{}"
	}

	return strings.Join([]string{"DdmGroupInfo", string(data)}, " ")
}
