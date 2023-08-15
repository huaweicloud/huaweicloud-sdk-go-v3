package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionAsMetricRequest Request Object
type ListFunctionAsMetricRequest struct {

	// 指标类型，默认值为failcount。
	Type *string `json:"type,omitempty"`

	// 起始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 指标类型，默认值为failcount。
	Limit *string `json:"limit,omitempty"`
}

func (o ListFunctionAsMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsMetricRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionAsMetricRequest", string(data)}, " ")
}
