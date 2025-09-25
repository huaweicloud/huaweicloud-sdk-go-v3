package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeMonthlyOperationReportTipStatusResponse Response Object
type ChangeMonthlyOperationReportTipStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeMonthlyOperationReportTipStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMonthlyOperationReportTipStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeMonthlyOperationReportTipStatusResponse", string(data)}, " ")
}
