package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateEndpointIpaddressRequest Request Object
type DisassociateEndpointIpaddressRequest struct {

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`

	// IP地址ID。
	IpaddressId string `json:"ipaddress_id"`
}

func (o DisassociateEndpointIpaddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEndpointIpaddressRequest struct{}"
	}

	return strings.Join([]string{"DisassociateEndpointIpaddressRequest", string(data)}, " ")
}
