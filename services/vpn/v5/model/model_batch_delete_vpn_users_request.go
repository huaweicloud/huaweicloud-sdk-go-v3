package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteVpnUsersRequest Request Object
type BatchDeleteVpnUsersRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	Body *BatchDeleteVpnUsersRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteVpnUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpnUsersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpnUsersRequest", string(data)}, " ")
}
