package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnServersByVgwRequest Request Object
type ListVpnServersByVgwRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`
}

func (o ListVpnServersByVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnServersByVgwRequest struct{}"
	}

	return strings.Join([]string{"ListVpnServersByVgwRequest", string(data)}, " ")
}
