package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesConfigurationsResponse Response Object
type ListInstancesConfigurationsResponse struct {

	// **参数解释**：  实例列表  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Entities *[]ApplicableInstanceInfo `json:"entities,omitempty"`

	// **参数解释**：  实例的限制数量  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceCountLimit *int32 `json:"instance_count_limit,omitempty"`

	// **参数解释**：  应用参数模板的实例总数  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesConfigurationsResponse", string(data)}, " ")
}
