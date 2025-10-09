package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessComputeAbilityPolicy **参数解释**：  设置serverless算力配置策略请求体。  **约束限制**：  不涉及。
type UpdateServerlessComputeAbilityPolicy struct {

	// **参数解释**：  单节点VCPUs伸缩下限。  **约束限制**：  不涉及。  **取值范围**：  取值范围可根据[查询数据库规格](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlFlavors.html)接口获取。  **默认取值**：  不涉及。
	MinVcpus int32 `json:"min_vcpus"`

	// **参数解释**：  单节点VCPUs伸缩上限。  **约束限制**：  不涉及。  **取值范围**：  ≥4。取值范围可根据[查询数据库规格](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlFlavors.html)接口获取。  **默认取值**：  不涉及。
	MaxVcpus int32 `json:"max_vcpus"`

	// **参数解释**：  节点算力同步，修改算力范围的同时，是否将小于最小算力的节点的当前算力同步至最小算力。  **约束限制**：  不涉及。  **取值范围**： - true: 节点算力同步。 - false: 节点算力不同步。  **默认取值**： false。
	NeedUpdateNodesComputeAbility *bool `json:"need_update_nodes_compute_ability,omitempty"`

	// **参数解释**：  是否增删只读节点。  **约束限制**：  - 存在数据库代理时，才可以使用增删只读节点功能。  - 使用增删节点功能时，避免使用读内网地址连接应用。  - 打开增删只读节点后，数据库代理的路由模式会变为负载均衡模式。  **取值范围**： - true: 开启增删只读节点。 - false: 不开启增删只读节点。  **默认取值**：  false。
	ScaleOutSwitch *bool `json:"scale_out_switch,omitempty"`

	// **参数解释**：  只读节点数量上限。  **约束限制**：  开启增删只读节点时才会生效, 且开启增删只读节点时该参数必选。  **取值范围**：  2-7。  **默认取值**：  不涉及。
	MaxReadonlyNodeCount *int32 `json:"max_readonly_node_count,omitempty"`

	// **参数解释**：  只读节点数量上限。  **约束限制**：  开启增删只读节点时才会生效, 且开启增删只读节点时该参数必选。  **取值范围**：  1-6。  **默认取值**：  不涉及。
	MinReadonlyNodeCount *int32 `json:"min_readonly_node_count,omitempty"`
}

func (o UpdateServerlessComputeAbilityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessComputeAbilityPolicy struct{}"
	}

	return strings.Join([]string{"UpdateServerlessComputeAbilityPolicy", string(data)}, " ")
}
