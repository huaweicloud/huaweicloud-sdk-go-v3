package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetInstancesOpsMetricNamesResponse Response Object
type GetInstancesOpsMetricNamesResponse struct {

	// **参数解释**：  监控指标项列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Metrics *[]MetricItem `json:"metrics,omitempty"`

	// **参数解释**：  CES命名空间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：  监控维度类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Dim            *string `json:"dim,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetInstancesOpsMetricNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesOpsMetricNamesResponse struct{}"
	}

	return strings.Join([]string{"GetInstancesOpsMetricNamesResponse", string(data)}, " ")
}
