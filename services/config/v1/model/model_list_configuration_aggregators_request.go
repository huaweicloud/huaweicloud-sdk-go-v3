package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationAggregatorsRequest Request Object
type ListConfigurationAggregatorsRequest struct {

	// 资源聚合器名称。
	AggregatorName *string `json:"aggregator_name,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListConfigurationAggregatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationAggregatorsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationAggregatorsRequest", string(data)}, " ")
}
