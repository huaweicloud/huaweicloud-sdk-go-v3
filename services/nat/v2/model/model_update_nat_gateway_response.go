package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty" xml:"nat_gateway"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayResponse", string(data)}, " ")
}
