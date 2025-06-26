package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanStageIdReq **参数解释**： 资源管理计划阶段对象。 **取值范围**： 不涉及。
type WorkloadPlanStageIdReq struct {

	// **参数解释**： 资源管理计划阶段ID。 **取值范围**： 不涉及。
	StageId string `json:"stage_id"`
}

func (o WorkloadPlanStageIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanStageIdReq struct{}"
	}

	return strings.Join([]string{"WorkloadPlanStageIdReq", string(data)}, " ")
}
