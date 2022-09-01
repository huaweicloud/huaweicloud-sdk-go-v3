package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty" xml:"nat_gateway"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayResponse", string(data)}, " ")
}
