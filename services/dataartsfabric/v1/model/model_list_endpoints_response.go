package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsResponse Response Object
type ListEndpointsResponse struct {

	// 符合条件的endpoint总数
	Total *int32 `json:"total,omitempty"`

	// 符合条件的Endpoint简要信息列表
	Endpoints *[]EndpointBriefInfo `json:"endpoints,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsResponse", string(data)}, " ")
}
