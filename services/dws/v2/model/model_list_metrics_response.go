package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// 响应码
	Code *int32 `json:"code,omitempty"`

	// 响应信息
	Msg *string `json:"msg,omitempty"`

	// 指标列表。
	Data *[]ClusterMetric `json:"data,omitempty"`

	// 总列表大小。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
