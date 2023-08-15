package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationAggregatorSourcesStatusResponse Response Object
type ShowConfigurationAggregatorSourcesStatusResponse struct {

	// 资源聚合器状态列表。
	AggregatedSourceStatuses *[]AggregatedSourceStatus `json:"aggregated_source_statuses,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowConfigurationAggregatorSourcesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAggregatorSourcesStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAggregatorSourcesStatusResponse", string(data)}, " ")
}
