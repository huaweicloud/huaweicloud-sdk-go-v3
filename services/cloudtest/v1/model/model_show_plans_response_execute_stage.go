package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划执行阶段信息
type ShowPlansResponseExecuteStage struct {
	// 缺陷个数

	DefectCount *int32 `json:"defect_count,omitempty"`
	// 已完成缺陷个数

	CompletedDefectCount *int32 `json:"completed_defect_count,omitempty"`
	// 用例通过率,按用例结果计算

	CasePassRate *string `json:"case_pass_rate,omitempty"`
	// 已执行用例数, 按用例状态统计

	ExecutedCaseCount *int32 `json:"executed_case_count,omitempty"`
}

func (o ShowPlansResponseExecuteStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlansResponseExecuteStage struct{}"
	}

	return strings.Join([]string{"ShowPlansResponseExecuteStage", string(data)}, " ")
}
