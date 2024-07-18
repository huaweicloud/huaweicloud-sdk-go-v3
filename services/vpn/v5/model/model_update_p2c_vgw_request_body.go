package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateP2cVgwRequestBody struct {
	P2cVpnGateway *UpdateP2cVgwRequestBodyContent `json:"p2c_vpn_gateway"`
}

func (o UpdateP2cVgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateP2cVgwRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateP2cVgwRequestBody", string(data)}, " ")
}
