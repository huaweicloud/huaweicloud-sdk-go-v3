package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModel 物理池创建请求体。
type PoolSpecModel struct {

	// **参数解释**：资源池类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Dedicate：物理资源池，独立的网络，支持网络打通，定制驱动，定制作业类型。 - Logical：逻辑资源池。没有独立的网络，不支持网络打通，资源池创建和扩缩容相较物理资源池更快。 **默认取值**：不涉及。
	Type string `json:"type"`

	// **参数解释**：资源池支持的作业类型。
	Scope []string `json:"scope"`

	// **参数解释**：资源池中的资源规格信列表，包括资源规格和相应规格的资源数量。
	Resources []PoolSpecModelResources `json:"resources"`

	Containernetwork *PoolSpecModelContainernetwork `json:"containernetwork,omitempty"`

	Network *PoolSpecModelNetwork `json:"network,omitempty"`

	// **参数解释**：资源池支持的作业规格列表。参数为作业规格名称。
	JobFlavors *[]string `json:"jobFlavors,omitempty"`

	Driver *PoolDriver `json:"driver,omitempty"`

	// **参数解释**：资源池的受限状态。状态可以叠加，比如9代表转包周期受限和冻结状态。 **取值范围**：可选值如下： - 0：代表不受限 - 1：转包周期受限 - 2：规格变更受限 - 4：服务受限 - 8：冻结 - 16：公安冻结（不可退订）
	ControlMode *int32 `json:"controlMode,omitempty"`
}

func (o PoolSpecModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModel struct{}"
	}

	return strings.Join([]string{"PoolSpecModel", string(data)}, " ")
}
