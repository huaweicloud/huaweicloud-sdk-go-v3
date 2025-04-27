package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsResponse Response Object
type ListEndpointsResponse struct {

	// 查询终端节点列表响应。
	Endpoints *[]EndpointResp `json:"endpoints,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsResponse", string(data)}, " ")
}
