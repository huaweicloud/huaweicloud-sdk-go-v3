package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointIpaddressesResponse Response Object
type ListEndpointIpaddressesResponse struct {

	// 列表数据结构。
	Ipaddresses    *[]IpaddressesData `json:"ipaddresses,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListEndpointIpaddressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointIpaddressesResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointIpaddressesResponse", string(data)}, " ")
}
