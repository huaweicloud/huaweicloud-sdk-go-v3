package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSharedResourcesResponse Response Object
type SearchSharedResourcesResponse struct {

	// 共享资源的信息列表。
	SharedResources *[]SharedResource `json:"shared_resources,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchSharedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedResourcesResponse struct{}"
	}

	return strings.Join([]string{"SearchSharedResourcesResponse", string(data)}, " ")
}
