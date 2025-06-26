package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceInternalEndpointsResponse Response Object
type ListInstanceInternalEndpointsResponse struct {

	// 内网访问列表
	InternalEndpoints *[]InternalEndpoint `json:"internal_endpoints,omitempty"`

	// 内网访问总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceInternalEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceInternalEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceInternalEndpointsResponse", string(data)}, " ")
}
