package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSummaryVo 执行阶段信息汇总
type ExecuteSummaryVo struct {

	// 已执行用例数
	ExecuteCaseNum *int32 `json:"execute_case_num,omitempty"`

	// 缺陷总数
	DefectNum *int32 `json:"defect_num,omitempty"`

	// 已完成缺陷数
	CompletedDefectNum *int32 `json:"completed_defect_num,omitempty"`

	// 测试用例通过率
	CaseSuccessRate *string `json:"case_success_rate,omitempty"`

	// 用例执行率
	CaseExecutionRate *string `json:"case_execution_rate,omitempty"`
}

func (o ExecuteSummaryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSummaryVo struct{}"
	}

	return strings.Join([]string{"ExecuteSummaryVo", string(data)}, " ")
}
