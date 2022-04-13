package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMonitoringDataResponse struct {
	// 容灾监控数据响应体集合

	Results *[]QueryDataGuardMonitorAndChartResp `json:"results,omitempty"`
	// 查询总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMonitoringDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitoringDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitoringDataResponse", string(data)}, " ")
}
