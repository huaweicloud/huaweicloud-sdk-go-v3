package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEndpointIpaddressesRequest struct {

	// 待查询的ip地址列表所在的endpoint的ID。
	EndpointId string `json:"endpoint_id"`
}

func (o ListEndpointIpaddressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointIpaddressesRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointIpaddressesRequest", string(data)}, " ")
}
