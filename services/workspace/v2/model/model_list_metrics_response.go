package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// 查询指标响应体
	Body           *[]Metric `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
