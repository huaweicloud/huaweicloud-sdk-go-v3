package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVpnConnectionPeerConfigurationResponse Response Object
type ExportVpnConnectionPeerConfigurationResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 对端配置
	PeerConfig *string `json:"peer_config,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ExportVpnConnectionPeerConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVpnConnectionPeerConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ExportVpnConnectionPeerConfigurationResponse", string(data)}, " ")
}
