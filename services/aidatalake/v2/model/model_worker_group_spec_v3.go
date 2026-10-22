package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkerGroupSpecV3 **参数解释**：Ray worker group配置信息。
type WorkerGroupSpecV3 struct {

	// **参数解释**：Worker Group的名称。 **约束限制**：不涉及。 **取值范围**：长度0~128字母、数字、中划线的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：最小副本数。 **约束限制**：不涉及。 **取值范围**：0~10000。 **默认取值**：不涉及。
	MinReplicas int32 `json:"min_replicas"`

	// **参数解释**：最大副本数。 **约束限制**：不涉及。 **取值范围**：1~10000。 **默认取值**：不涉及。
	MaxReplicas int32 `json:"max_replicas"`

	Resource *ResourceSpecV3 `json:"resource"`
}

func (o WorkerGroupSpecV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerGroupSpecV3 struct{}"
	}

	return strings.Join([]string{"WorkerGroupSpecV3", string(data)}, " ")
}
