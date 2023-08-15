package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsResponse Response Object
type ListEndpointsResponse struct {

	// 端点详情
	Endpoints *[]EndpointObjResp `json:"endpoints,omitempty"`

	// 满足条件的端点个数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsResponse", string(data)}, " ")
}
