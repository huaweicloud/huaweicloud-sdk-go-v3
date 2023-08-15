package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricDataResponse Response Object
type ListMetricDataResponse struct {

	// 指标数据列表
	Datapoints     *[]MetricData `json:"datapoints,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ListMetricDataResponse", string(data)}, " ")
}
