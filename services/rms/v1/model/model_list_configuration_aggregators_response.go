package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigurationAggregatorsResponse struct {

	// 资源聚合器列表。
	ConfigurationAggregators *[]ConfigurationAggregatorResp `json:"configuration_aggregators,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConfigurationAggregatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationAggregatorsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationAggregatorsResponse", string(data)}, " ")
}
