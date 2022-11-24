package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateEndpointIpaddressRequest struct {

	// endpoint的ID。
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
