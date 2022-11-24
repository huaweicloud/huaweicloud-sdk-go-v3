package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayResponse", string(data)}, " ")
}
