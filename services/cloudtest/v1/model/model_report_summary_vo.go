package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportSummaryVo 报告阶段信息汇总
type ReportSummaryVo struct {

	// 用例通过率
	CaseSuccessRate *string `json:"case_success_rate,omitempty"`

	// 用例完成率
	CaseCompleteRate *string `json:"case_complete_rate,omitempty"`
}

func (o ReportSummaryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportSummaryVo struct{}"
	}

	return strings.Join([]string{"ReportSummaryVo", string(data)}, " ")
}
