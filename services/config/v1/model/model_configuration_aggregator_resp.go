package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationAggregatorResp 资源聚合器响应体。
type ConfigurationAggregatorResp struct {

	// 资源聚合器名称。
	AggregatorName *string `json:"aggregator_name,omitempty"`

	// 资源聚合器ID。
	AggregatorId *string `json:"aggregator_id,omitempty"`

	// 资源聚合器标识符。
	AggregatorUrn *string `json:"aggregator_urn,omitempty"`

	// 聚合器类型。
	AggregatorType *string `json:"aggregator_type,omitempty"`

	AccountAggregationSources *AccountAggregationSource `json:"account_aggregation_sources,omitempty"`

	// 资源聚合器更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 资源聚合器创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ConfigurationAggregatorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationAggregatorResp struct{}"
	}

	return strings.Join([]string{"ConfigurationAggregatorResp", string(data)}, " ")
}
