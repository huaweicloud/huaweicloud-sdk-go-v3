package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonthlyOperaReportNotifyInfoRequest Request Object
type ShowMonthlyOperaReportNotifyInfoRequest struct {
}

func (o ShowMonthlyOperaReportNotifyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthlyOperaReportNotifyInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowMonthlyOperaReportNotifyInfoRequest", string(data)}, " ")
}
