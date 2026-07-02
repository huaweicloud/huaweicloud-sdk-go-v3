package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVpnConnectionPeerConfigurationRequest Request Object
type ExportVpnConnectionPeerConfigurationRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`

	Body *ExportPeerConfigurationRequestBody `json:"body,omitempty"`
}

func (o ExportVpnConnectionPeerConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVpnConnectionPeerConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ExportVpnConnectionPeerConfigurationRequest", string(data)}, " ")
}
