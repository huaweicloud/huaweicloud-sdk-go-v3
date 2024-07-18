package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServerRequestBody struct {
	VpnServer *CreateServerRequest `json:"vpn_server,omitempty"`
}

func (o CreateServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateServerRequestBody", string(data)}, " ")
}
