package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Statistic **参数解释**： 资源数量详情。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type Statistic struct {

	// **参数解释**： 资源名称。 **约束限制**： 不涉及。 **取值范围**： - cluster.total：总集群（个）。 - cluster.normal：可用集群（个）。 - instance.total：总节点（个）。 - instance.normal：可用节点（个）。 - storage.total：总容量（GB）。  **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 资源数量值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value float64 `json:"value"`

	// **参数解释**： 资源数量单位。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Unit string `json:"unit"`
}

func (o Statistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Statistic struct{}"
	}

	return strings.Join([]string{"Statistic", string(data)}, " ")
}
