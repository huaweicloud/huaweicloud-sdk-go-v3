package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonthlyOperationReportDetailResponse Response Object
type ShowMonthlyOperationReportDetailResponse struct {
	OperationSummaryInfo *OperationSummaryInfo `json:"operation_summary_info,omitempty"`

	ProtectInfo *ProtectInfo `json:"protect_info,omitempty"`

	RiskHandleInfo *RiskHandleInfo `json:"risk_handle_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowMonthlyOperationReportDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthlyOperationReportDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMonthlyOperationReportDetailResponse", string(data)}, " ")
}
