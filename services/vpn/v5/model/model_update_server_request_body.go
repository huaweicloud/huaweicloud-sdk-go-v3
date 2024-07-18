package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServerRequestBody struct {
	VpnServer *UpdateServerRequest `json:"vpn_server,omitempty"`
}

func (o UpdateServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerRequestBody", string(data)}, " ")
}
