package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEndpointsResponse struct {

	// 终端节点列表。
	Endpoints *[]EndpointDetail `json:"endpoints,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsResponse", string(data)}, " ")
}
