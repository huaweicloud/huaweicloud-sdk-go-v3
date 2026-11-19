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

	// **参数解释**：  Scale to Zero开关。  **取值范围**： - true: 开启Scale to Zero。 - false: 关闭Scale to Zero。
	ScaleToZero *bool `json:"scale_to_zero,omitempty"`

	// **参数解释**：  空闲持续时长阈值，单位是秒。  所有节点活跃连接数为0且持续达到该阈值时触发Scale to Zero（停止计算计费）。  **取值范围**：  300-3600。
	IdleDurationSeconds *int32 `json:"idle_duration_seconds,omitempty"`

	// **参数解释**：  Scale to Zero开启后实例的运行状态。  **取值范围**： - ACTIVE：实例已唤醒，正常计费。 - SCALED_TO_ZERO：计算资源停止计费。
	ScaleToZeroStatus *string `json:"scale_to_zero_status,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowServerlessComputeAbilityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerlessComputeAbilityPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowServerlessComputeAbilityPolicyResponse", string(data)}, " ")
}
