package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceNetworkResponse Response Object
type ChangeInstanceNetworkResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 安全组
	SecurityGrpStatus *string `json:"security_grp_status,omitempty"`

	// 防火墙状态
	FirewallStatus *bool `json:"firewall_status,omitempty"`

	// 公共EIP状态
	PublicEipStatus *bool `json:"public_eip_status,omitempty"`

	// 防火墙状态(兼容)
	Nics *bool `json:"nics,omitempty"`

	// 公共EIP状态(兼容)
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
