package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecUpdate 资源池描述信息更新参数。
type PoolSpecUpdate struct {

	// **参数解释**：更新启用的作业类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Train：训练作业 - Infer：推理作业 - Notebook：Notebook作业 **默认取值**：不涉及。
	Scope *[]string `json:"scope,omitempty"`

	// **参数解释**：更新的资源规格列表。
	Resources *[]PoolSpecUpdateResources `json:"resources,omitempty"`

	// **参数解释**：资源池支持的作业规格信息列表。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobFlavors *[]string `json:"jobFlavors,omitempty"`

	Driver *PoolDriver `json:"driver,omitempty"`
}

func (o PoolSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecUpdate struct{}"
	}

	return strings.Join([]string{"PoolSpecUpdate", string(data)}, " ")
}
