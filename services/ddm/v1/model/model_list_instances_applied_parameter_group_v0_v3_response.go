package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesAppliedParameterGroupV0V3Response Response Object
type ListInstancesAppliedParameterGroupV0V3Response struct {

	// **参数解释**：  查询可应用的实例列表返回相关信息的集合。  **参数范围**：  不涉及。
	Entities *[]ApplicableInstance `json:"entities,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **约束限制**：  不涉及。  **取值范围**：  大于0且小于等于128。  **默认取值**：  默认值是10。
	InstanceCountLimit *int32 `json:"instance_count_limit,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ListInstancesAppliedParameterGroupV0V3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesAppliedParameterGroupV0V3Response struct{}"
	}

	return strings.Join([]string{"ListInstancesAppliedParameterGroupV0V3Response", string(data)}, " ")
}
