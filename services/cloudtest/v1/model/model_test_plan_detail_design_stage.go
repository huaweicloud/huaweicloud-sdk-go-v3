package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestPlanDetailDesignStage 测试计划设计阶段信息
type TestPlanDetailDesignStage struct {

	// 用例个数
	CaseCount *int32 `json:"case_count,omitempty"`

	// 需求个数
	IssueCount *int32 `json:"issue_count,omitempty"`

	// 已被用例关联的需求个数
	IssueCoveredCount *string `json:"issue_covered_count,omitempty"`
}

func (o TestPlanDetailDesignStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailDesignStage struct{}"
	}

	return strings.Join([]string{"TestPlanDetailDesignStage", string(data)}, " ")
}
