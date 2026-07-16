package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecCreation 资源池创建请求体。
type PoolSpecCreation struct {

	// **参数解释**：资源池类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Dedicate：物理资源池，独立的网络，支持网络打通，定制驱动，定制作业类型 - Logical：逻辑资源池。没有独立的网络，不支持网络打通，资源池创建和扩缩容相较物理资源池更快。 **默认取值**：不涉及。
	Type string `json:"type"`

	// **参数解释**：资源池支持的作业类型。 **约束限制**：不涉及。 **取值范围**：用户创建标准资源池时至少选择一种，物理资源池支持全部选择。可选值如下： - Train：训练作业 - Infer：推理作业 - Notebook：Notebook作业 **默认取值**：不涉及。
	Scope []string `json:"scope"`

	// **参数解释**：资源池中的资源规格信列表，包括资源规格和相应规格的资源数量。 **约束限制**：不涉及。
	Resources []PoolResourceFlavor `json:"resources"`

	Network *PoolSpecCreationNetwork `json:"network"`

	// **参数解释**：资源池支持的作业规格信息列表，内容为作业规格名称。 **约束限制**：不涉及。
	JobFlavors *[]string `json:"jobFlavors,omitempty"`

	Driver *PoolDriver `json:"driver,omitempty"`
}

func (o PoolSpecCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecCreation struct{}"
	}

	return strings.Join([]string{"PoolSpecCreation", string(data)}, " ")
}
