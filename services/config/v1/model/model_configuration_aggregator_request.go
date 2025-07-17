package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationAggregatorRequest 资源聚合器请求体。
type ConfigurationAggregatorRequest struct {

	// 资源聚合器名称。
	AggregatorName string `json:"aggregator_name"`

	// 聚合器类型（ACCOUNT | ORGANIZATION）。
	AggregatorType string `json:"aggregator_type"`

	AccountAggregationSources *AccountAggregationSource `json:"account_aggregation_sources,omitempty"`

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ConfigurationAggregatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationAggregatorRequest struct{}"
	}

	return strings.Join([]string{"ConfigurationAggregatorRequest", string(data)}, " ")
}
