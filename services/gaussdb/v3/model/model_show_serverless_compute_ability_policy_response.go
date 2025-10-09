package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerlessComputeAbilityPolicyResponse Response Object
type ShowServerlessComputeAbilityPolicyResponse struct {

	// **参数描述**：  当前算力。  **取值范围**：  介于最大算力和最小算力之间。
	CurrentVcpus *string `json:"current_vcpus,omitempty"`

	// **参数解释**：  最小算力。  **取值范围**：  ≥0.5。
	MinVcpus *string `json:"min_vcpus,omitempty"`

	// **参数解释**：  最大算力。  **取值范围**：  ≥4。
	MaxVcpus *string `json:"max_vcpus,omitempty"`

	// **参数解释**:  增删只读节点开关。  **取值范围**： - true: 开启增删只读节点。 - false: 关闭增删只读节点。
	ScaleOutSwitch *bool `json:"scale_out_switch,omitempty"`

	// **参数解释**：  最大只读节点个数。  **取值范围**：  ≥2。
	MaxReadonlyNodeCount *int32 `json:"max_readonly_node_count,omitempty"`

	// **参数解释**：  最小只读节点个数。  **取值范围**：  ≥1。
	MinReadonlyNodeCount *int32 `json:"min_readonly_node_count,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowServerlessComputeAbilityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerlessComputeAbilityPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowServerlessComputeAbilityPolicyResponse", string(data)}, " ")
}
