package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeInstanceNetworkResponse struct {

	// 状态
	Type *string `json:"type,omitempty"`

	// 安全组
	SecurityGroups *string `json:"security_groups,omitempty"`

	// 防火墙状态
	FirewallStatus *bool `json:"firewall_status,omitempty"`

	// 公共EIP状态
	PublicEipStatu *bool `json:"public_eip_statu,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ChangeInstanceNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceNetworkResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceNetworkResponse", string(data)}, " ")
}
