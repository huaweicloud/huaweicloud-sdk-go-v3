package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewayResponse Response Object
type CreateNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayResponse", string(data)}, " ")
}
