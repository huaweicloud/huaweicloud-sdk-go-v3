package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTypesResponse Response Object
type ListResourceTypesResponse struct {

	// 云服务资源的详细信息列表。
	ResourceTypes *[]ResourceTypesSummary `json:"resource_types,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResourceTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTypesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceTypesResponse", string(data)}, " ")
}
