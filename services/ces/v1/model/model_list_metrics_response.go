package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// 指标信息列表
	Metrics *[]MetricInfoList `json:"metrics,omitempty"`

	MetaData       *MetaData `json:"meta_data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
