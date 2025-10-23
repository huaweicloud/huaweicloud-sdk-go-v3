package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// 指标记录总条数
	Count *int64 `json:"count,omitempty"`

	// 指标数据
	Datapoints     *[]ListMetricsResponseBodyDatapoints `json:"datapoints,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
