package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateEndpointIpaddressResponse Response Object
type DisassociateEndpointIpaddressResponse struct {
	Endpoint       *EndpointResp `json:"endpoint,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DisassociateEndpointIpaddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEndpointIpaddressResponse struct{}"
	}

	return strings.Join([]string{"DisassociateEndpointIpaddressResponse", string(data)}, " ")
}
