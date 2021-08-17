package model

import (
	"encoding/json"

	"strings"
)

// 测试计划设计阶段信息
type ShowPlansResponseDesignStage struct {
	// 用例个数

	CaseCount *int32 `json:"case_count,omitempty"`
	// 需求个数

	IssueCount *int32 `json:"issue_count,omitempty"`
	// 已被用例关联的需求个数

	IssueCoveredCount *string `json:"issue_covered_count,omitempty"`
}

func (o ShowPlansResponseDesignStage) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPlansResponseDesignStage struct{}"
	}

	return strings.Join([]string{"ShowPlansResponseDesignStage", string(data)}, " ")
}
