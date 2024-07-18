package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportClientConfigRequest Request Object
type ExportClientConfigRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`
}

func (o ExportClientConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportClientConfigRequest struct{}"
	}

	return strings.Join([]string{"ExportClientConfigRequest", string(data)}, " ")
}
