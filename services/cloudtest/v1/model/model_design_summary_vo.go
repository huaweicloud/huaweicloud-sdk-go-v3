package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesignSummaryVo 设计阶段信息汇总
type DesignSummaryVo struct {

	// 需求总数
	IssueNum *int32 `json:"issue_num,omitempty"`

	// 已覆盖需求数
	IssueCoverNum *int32 `json:"issue_cover_num,omitempty"`

	// 用例数
	CaseNum *int32 `json:"case_num,omitempty"`
}

func (o DesignSummaryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesignSummaryVo struct{}"
	}

	return strings.Join([]string{"DesignSummaryVo", string(data)}, " ")
}
