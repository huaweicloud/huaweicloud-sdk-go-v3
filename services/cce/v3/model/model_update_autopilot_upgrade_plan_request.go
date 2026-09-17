package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutopilotUpgradePlanRequest Request Object
type UpdateAutopilotUpgradePlanRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释：** 集群自动升级计划ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpgradePlanId string `json:"upgrade_plan_id"`

	Body *DelayUpgradePlanRequestBody `json:"body,omitempty"`
}

func (o UpdateAutopilotUpgradePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutopilotUpgradePlanRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutopilotUpgradePlanRequest", string(data)}, " ")
}
