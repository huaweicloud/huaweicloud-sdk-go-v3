package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonthlyOperationReportDetailRequest Request Object
type ShowMonthlyOperationReportDetailRequest struct {

	// **参数解释**: reportId，生成的月度运营报告年月 **约束限制**: 不涉及 **取值范围**: 字符长度1-32位 **默认取值**: 不涉及
	ReportId string `json:"report_id"`
}

func (o ShowMonthlyOperationReportDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthlyOperationReportDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMonthlyOperationReportDetailRequest", string(data)}, " ")
}
