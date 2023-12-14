package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrackedResourcesResponse Response Object
type ListTrackedResourcesResponse struct {

	// 资源列表
	Resources *[]ResourceEntity `json:"resources,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTrackedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListTrackedResourcesResponse", string(data)}, " ")
}
