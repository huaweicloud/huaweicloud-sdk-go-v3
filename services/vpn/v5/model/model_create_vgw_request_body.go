package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVgwRequestBody struct {
	VpnGateway *CreateVgwRequestBodyContent `json:"vpn_gateway"`
}

func (o CreateVgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVgwRequestBody", string(data)}, " ")
}
