package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServerResponseBodyVpnServer struct {

	// VPN服务端 ID
	Id *string `json:"id,omitempty"`
}

func (o CreateServerResponseBodyVpnServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerResponseBodyVpnServer struct{}"
	}

	return strings.Join([]string{"CreateServerResponseBodyVpnServer", string(data)}, " ")
}
