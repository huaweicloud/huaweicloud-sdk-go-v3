package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointIpaddressesRequest Request Object
type ListEndpointIpaddressesRequest struct {

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`
}

func (o ListEndpointIpaddressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointIpaddressesRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointIpaddressesRequest", string(data)}, " ")
}
