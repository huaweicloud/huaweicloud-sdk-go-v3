package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterPlanBo **参数解释**： 更新逻辑集群增删计划请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type UpdateLogicalClusterPlanBo struct {

	// **参数解释**： 更新逻辑集群增删计划细节信息列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Actions []UpdateLogicalClusterPlanActions `json:"actions"`
}

func (o UpdateLogicalClusterPlanBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterPlanBo struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterPlanBo", string(data)}, " ")
}
