package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMoniterMetricStatsResponse Response Object
type ShowMoniterMetricStatsResponse struct {

	// 数据点
	Result         *[]Result `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowMoniterMetricStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMoniterMetricStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowMoniterMetricStatsResponse", string(data)}, " ")
}
