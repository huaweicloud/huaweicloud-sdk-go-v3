package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAllResourcesResponse struct {
	// 资源列表

	Resources *[]ResourceEntity `json:"resources,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListAllResourcesResponse", string(data)}, " ")
}
