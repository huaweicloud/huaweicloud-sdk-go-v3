package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchDistinctSharedResourcesResponse struct {

	// 不同资源的信息列表。
	DistinctSharedResources *[]DistinctSharedResource `json:"distinct_shared_resources,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchDistinctSharedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctSharedResourcesResponse struct{}"
	}

	return strings.Join([]string{"SearchDistinctSharedResourcesResponse", string(data)}, " ")
}
