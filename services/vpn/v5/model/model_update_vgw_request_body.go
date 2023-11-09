package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVgwRequestBody struct {
	VpnGateway *UpdateVgwRequestBodyContent `json:"vpn_gateway"`
}

func (o UpdateVgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVgwRequestBody", string(data)}, " ")
}
