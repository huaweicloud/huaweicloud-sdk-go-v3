package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAggregateDiscoveredResourcesResponse struct {

	// 资源信息列表。
	ResourceIdentifiers *[]ResourceIdentifier `json:"resource_identifiers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAggregateDiscoveredResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregateDiscoveredResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListAggregateDiscoveredResourcesResponse", string(data)}, " ")
}
