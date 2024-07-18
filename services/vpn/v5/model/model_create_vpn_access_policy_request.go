package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnAccessPolicyRequest Request Object
type CreateVpnAccessPolicyRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateVpnAccessPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateVpnAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateVpnAccessPolicyRequest", string(data)}, " ")
}
