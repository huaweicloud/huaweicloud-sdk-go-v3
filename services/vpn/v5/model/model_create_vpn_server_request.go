package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnServerRequest Request Object
type CreateVpnServerRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateServerRequestBody `json:"body,omitempty"`
}

func (o CreateVpnServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnServerRequest struct{}"
	}

	return strings.Join([]string{"CreateVpnServerRequest", string(data)}, " ")
}
