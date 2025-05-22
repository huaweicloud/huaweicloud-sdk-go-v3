package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchPlanStageRequest Request Object
type SwitchPlanStageRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 计划ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanId string `json:"plan_id"`

	Body *WorkloadPlanStageIdReq `json:"body,omitempty"`
}

func (o SwitchPlanStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchPlanStageRequest struct{}"
	}

	return strings.Join([]string{"SwitchPlanStageRequest", string(data)}, " ")
}
