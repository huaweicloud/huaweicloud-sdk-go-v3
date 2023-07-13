package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInstanceMetricDataResponse Response Object
type ShowTaskInstanceMetricDataResponse struct {

	// 监控数据列表
	DataPoints     *[]DataPointDto `json:"data_points,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTaskInstanceMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstanceMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInstanceMetricDataResponse", string(data)}, " ")
}
