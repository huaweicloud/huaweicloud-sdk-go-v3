package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExceptionMetricsResponse Response Object
type ListExceptionMetricsResponse struct {

	// 指标数据列表
	Metrics        *[]ExceptionMetricData `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListExceptionMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExceptionMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListExceptionMetricsResponse", string(data)}, " ")
}
