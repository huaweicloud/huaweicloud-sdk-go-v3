package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionAsMetricResponse Response Object
type ListFunctionAsMetricResponse struct {

	// 函数Urn及其指标数
	Values *[]FunctionMetric `json:"values,omitempty"`

	// 下次读取位置
	NextMarker *int64 `json:"next_marker,omitempty"`

	// 返回函数总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionAsMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsMetricResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionAsMetricResponse", string(data)}, " ")
}
