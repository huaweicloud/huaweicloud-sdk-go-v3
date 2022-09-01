package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划执行阶段信息
type TestPlanDetailExecuteStage struct {

	// 缺陷个数
	DefectCount *int32 `json:"defect_count,omitempty" xml:"defect_count"`

	// 已完成缺陷个数
	CompletedDefectCount *int32 `json:"completed_defect_count,omitempty" xml:"completed_defect_count"`

	// 用例通过率,按用例结果计算
	CasePassRate *string `json:"case_pass_rate,omitempty" xml:"case_pass_rate"`

	// 已执行用例数, 按用例状态统计
	ExecutedCaseCount *int32 `json:"executed_case_count,omitempty" xml:"executed_case_count"`
}

func (o TestPlanDetailExecuteStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailExecuteStage struct{}"
	}

	return strings.Join([]string{"TestPlanDetailExecuteStage", string(data)}, " ")
}
