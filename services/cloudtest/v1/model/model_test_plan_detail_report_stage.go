package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestPlanDetailReportStage 测试计划完成阶段信息
type TestPlanDetailReportStage struct {

	// 用例完成率,按状态统计
	CaseCompleteRate *string `json:"case_complete_rate,omitempty"`
}

func (o TestPlanDetailReportStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailReportStage struct{}"
	}

	return strings.Join([]string{"TestPlanDetailReportStage", string(data)}, " ")
}
