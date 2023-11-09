package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnConnectionRequestBodyContent struct {

	// VPN连接名称
	Name *string `json:"name,omitempty"`

	// 对端网关ID
	CgwId *string `json:"cgw_id,omitempty"`

	// 对端网段
	PeerSubnets *[]string `json:"peer_subnets,omitempty"`

	// 本端隧道口地址
	TunnelLocalAddress *string `json:"tunnel_local_address,omitempty"`

	// 对端隧道口地址
	TunnelPeerAddress *string `json:"tunnel_peer_address,omitempty"`

	// 预共享密钥，只能包含大写字母、小写字母、数字和特殊字符(~!@#$%^()-_+={ },./:;)且至少包含四种字符的三种
	Psk *string `json:"psk,omitempty"`

	// 策略模式的策略规则组
	PolicyRules *[]PolicyRule `json:"policy_rules,omitempty"`

	Ikepolicy *UpdateIkePolicy `json:"ikepolicy,omitempty"`

	Ipsecpolicy *UpdateIpsecPolicy `json:"ipsecpolicy,omitempty"`
}

func (o UpdateVpnConnectionRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionRequestBodyContent", string(data)}, " ")
}
