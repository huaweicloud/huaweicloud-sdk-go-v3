package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNetworkConfigurationResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 安全组状态
	SecurityGrpStatus *bool `json:"security_grp_status,omitempty"`

	// 防火墙状态
	FirewallStatus *bool `json:"firewall_status,omitempty"`

	// 公网IP状态
	PublicEipStatus *bool `json:"public_eip_status,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ShowNetworkConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowNetworkConfigurationResponse", string(data)}, " ")
}
