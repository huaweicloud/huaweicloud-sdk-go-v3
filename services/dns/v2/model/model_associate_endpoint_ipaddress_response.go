package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateEndpointIpaddressResponse Response Object
type AssociateEndpointIpaddressResponse struct {
	Endpoint       *EndpointResp `json:"endpoint,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o AssociateEndpointIpaddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateEndpointIpaddressResponse struct{}"
	}

	return strings.Join([]string{"AssociateEndpointIpaddressResponse", string(data)}, " ")
}
