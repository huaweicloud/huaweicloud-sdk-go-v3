package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateP2cVgwRequestBody struct {
	P2cVpnGateway *CreateP2cVgwRequestBodyContent `json:"p2c_vpn_gateway"`
}

func (o CreateP2cVgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2cVgwRequestBody struct{}"
	}

	return strings.Join([]string{"CreateP2cVgwRequestBody", string(data)}, " ")
}
