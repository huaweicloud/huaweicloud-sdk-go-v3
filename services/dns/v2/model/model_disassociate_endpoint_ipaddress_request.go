package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateEndpointIpaddressRequest struct {

	// 待解绑定的ip地址所属endpoint的ID。
	EndpointId string `json:"endpoint_id"`

	// 待解绑定ip地址的ID。
	IpaddressId string `json:"ipaddress_id"`
}

func (o DisassociateEndpointIpaddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEndpointIpaddressRequest struct{}"
	}

	return strings.Join([]string{"DisassociateEndpointIpaddressRequest", string(data)}, " ")
}
