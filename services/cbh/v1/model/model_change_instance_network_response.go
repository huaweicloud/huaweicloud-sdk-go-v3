package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceNetworkResponse Response Object
type ChangeInstanceNetworkResponse struct {

	// 云堡垒机实例网络状态。下面3个正常则正常，有一个不正常，网络状态为失败。
	Status *bool `json:"status,omitempty"`

	// 云堡垒机实例安全组状态。 - true  正常 - false 失败
	SecurityGrpStatus *bool `json:"security_grp_status,omitempty"`

	// 云堡垒机实例防火墙状态。 - true  正常 - false 失败
	FirewallStatus *bool `json:"firewall_status,omitempty"`

	// 云堡垒机实例公网IP状态。 - true  正常 - false 失败
	PublicEipStatus *bool `json:"public_eip_status,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ChangeInstanceNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceNetworkResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceNetworkResponse", string(data)}, " ")
}
