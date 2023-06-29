package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNatGatewayResponse Response Object
type DeleteNatGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayResponse", string(data)}, " ")
}
