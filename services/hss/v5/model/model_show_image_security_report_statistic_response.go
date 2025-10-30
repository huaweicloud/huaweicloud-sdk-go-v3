package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageSecurityReportStatisticResponse Response Object
type ShowImageSecurityReportStatisticResponse struct {

	// 导出记录总数
	TotalNum       *int64 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageSecurityReportStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageSecurityReportStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowImageSecurityReportStatisticResponse", string(data)}, " ")
}
