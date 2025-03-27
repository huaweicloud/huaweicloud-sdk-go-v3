package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnConnectionsLogConfigResponse Response Object
type UpdateVpnConnectionsLogConfigResponse struct {
	LogConfig *ConnectionsLogConfig `json:"log_config,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnConnectionsLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionsLogConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionsLogConfigResponse", string(data)}, " ")
}
