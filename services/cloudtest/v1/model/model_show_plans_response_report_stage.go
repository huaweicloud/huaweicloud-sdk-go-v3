package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划完成阶段信息
type ShowPlansResponseReportStage struct {
	// 用例完成率,按状态统计

	CaseCompleteRate *string `json:"case_complete_rate,omitempty"`
}

func (o ShowPlansResponseReportStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlansResponseReportStage struct{}"
	}

	return strings.Join([]string{"ShowPlansResponseReportStage", string(data)}, " ")
}
