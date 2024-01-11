package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanStageIdReq 资源管理计划阶段ID
type WorkloadPlanStageIdReq struct {

	// 计划阶段ID
	StageId string `json:"stage_id"`
}

func (o WorkloadPlanStageIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanStageIdReq struct{}"
	}

	return strings.Join([]string{"WorkloadPlanStageIdReq", string(data)}, " ")
}
