package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnConnectionRequestBody struct {
	VpnConnection *UpdateVpnConnectionRequestBodyContent `json:"vpn_connection"`
}

func (o UpdateVpnConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionRequestBody", string(data)}, " ")
}
