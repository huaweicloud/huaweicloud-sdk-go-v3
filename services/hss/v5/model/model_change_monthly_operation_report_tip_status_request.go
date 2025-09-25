package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeMonthlyOperationReportTipStatusRequest Request Object
type ChangeMonthlyOperationReportTipStatusRequest struct {
}

func (o ChangeMonthlyOperationReportTipStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMonthlyOperationReportTipStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeMonthlyOperationReportTipStatusRequest", string(data)}, " ")
}
