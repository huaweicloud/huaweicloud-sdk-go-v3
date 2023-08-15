package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointGroupsResponse Response Object
type ListEndpointGroupsResponse struct {

	// 终端节点组列表。
	EndpointGroups *[]EndpointGroupDetail `json:"endpoint_groups,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointGroupsResponse", string(data)}, " ")
}
