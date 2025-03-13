package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateVpnUsersRequest Request Object
type BatchCreateVpnUsersRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *BatchCreateVpnUsersRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVpnUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnUsersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnUsersRequest", string(data)}, " ")
}
