package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateEndpointIpaddressRequest Request Object
type AssociateEndpointIpaddressRequest struct {

	// 终端节点ID。
	EndpointId string `json:"endpoint_id"`

	Body *Ipaddresses `json:"body,omitempty"`
}

func (o AssociateEndpointIpaddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateEndpointIpaddressRequest struct{}"
	}

	return strings.Join([]string{"AssociateEndpointIpaddressRequest", string(data)}, " ")
}
