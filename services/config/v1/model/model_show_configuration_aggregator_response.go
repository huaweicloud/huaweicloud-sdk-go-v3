package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationAggregatorResponse Response Object
type ShowConfigurationAggregatorResponse struct {

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
	CreatedAt      *string `json:"created_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfigurationAggregatorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAggregatorResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAggregatorResponse", string(data)}, " ")
}
