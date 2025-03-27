package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionsLogConfigResponse Response Object
type ShowVpnConnectionsLogConfigResponse struct {
	LogConfig *ConnectionsLogConfig `json:"log_config,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnConnectionsLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionsLogConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionsLogConfigResponse", string(data)}, " ")
}
