package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSummaryBodyResult 结果
type ShowReportSummaryBodyResult struct {
	Summary *ShowReportSummary `json:"summary,omitempty"`

	// 单元测试报告列表
	SubSummarys *[]ShowReportSummary `json:"sub_summarys,omitempty"`
}

func (o ShowReportSummaryBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSummaryBodyResult struct{}"
	}

	return strings.Join([]string{"ShowReportSummaryBodyResult", string(data)}, " ")
}
