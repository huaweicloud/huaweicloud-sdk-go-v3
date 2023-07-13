package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDailyReportResponse Response Object
type ListDailyReportResponse struct {

	// 24小时内的流量数据
	Data           *[]DailyData `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDailyReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDailyReportResponse struct{}"
	}

	return strings.Join([]string{"ListDailyReportResponse", string(data)}, " ")
}
