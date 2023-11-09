package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnConnectionRequestBody struct {
	VpnConnection *CreateVpnConnectionRequestBodyContent `json:"vpn_connection"`
}

func (o CreateVpnConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpnConnectionRequestBody", string(data)}, " ")
}
