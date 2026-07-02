package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateVpnConnectionRequestBodyContent struct {

	// VPN连接名称
	Name *string `json:"name,omitempty"`

	// VPN网关ID
	VgwId string `json:"vgw_id"`

	// VGW IP
	VgwIp string `json:"vgw_ip"`

	// 连接模式 允许范围[policy, static, bgp] policy: 策略模式 static: 静态路由模式 bgp: bgp路由模式
	Style *CreateVpnConnectionRequestBodyContentStyle `json:"style,omitempty"`

	// 对端网关ID
	CgwId string `json:"cgw_id"`

	// 对端子网
	PeerSubnets *[]string `json:"peer_subnets,omitempty"`

	// 本端隧道口地址
	TunnelLocalAddress *string `json:"tunnel_local_address,omitempty"`

	// 对端隧道口地址
	TunnelPeerAddress *string `json:"tunnel_peer_address,omitempty"`

	// 开启NQA检测
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 开启分支互联
	EnableHub *bool `json:"enable_hub,omitempty"`

	// 开启健康检查
	EnableHealthCheck *bool `json:"enable_health_check,omitempty"`

	// 预共享密钥，只能包含大写字母、小写字母、数字和特殊字符(~!@#$%^()-_+={ },./:;)且至少包含四种字符的三种
	Psk *string `json:"psk,omitempty"`

	// 策略模式的策略规则组
	PolicyRules *[]PolicyRule `json:"policy_rules,omitempty"`

	Ikepolicy *IkePolicy `json:"ikepolicy,omitempty"`

	Ipsecpolicy *IpsecPolicy `json:"ipsecpolicy,omitempty"`

	// 连接的HA角色
	HaRole *CreateVpnConnectionRequestBodyContentHaRole `json:"ha_role,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`

	// 使能ipv6的对端子网
	PeerSubnetsV6 *[]string `json:"peer_subnets_v6,omitempty"`

	// 策略模式的ipv6策略规则组
	PolicyRulesV6 *[]PolicyRule `json:"policy_rules_v6,omitempty"`
}

func (o CreateVpnConnectionRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnConnectionRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateVpnConnectionRequestBodyContent", string(data)}, " ")
}

type CreateVpnConnectionRequestBodyContentStyle struct {
	value string
}

type CreateVpnConnectionRequestBodyContentStyleEnum struct {
	POLICY CreateVpnConnectionRequestBodyContentStyle
	STATIC CreateVpnConnectionRequestBodyContentStyle
	BGP    CreateVpnConnectionRequestBodyContentStyle
}

func GetCreateVpnConnectionRequestBodyContentStyleEnum() CreateVpnConnectionRequestBodyContentStyleEnum {
	return CreateVpnConnectionRequestBodyContentStyleEnum{
		POLICY: CreateVpnConnectionRequestBodyContentStyle{
			value: "policy",
		},
		STATIC: CreateVpnConnectionRequestBodyContentStyle{
			value: "static",
		},
		BGP: CreateVpnConnectionRequestBodyContentStyle{
			value: "bgp",
		},
	}
}

func (c CreateVpnConnectionRequestBodyContentStyle) Value() string {
	return c.value
}

func (c CreateVpnConnectionRequestBodyContentStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpnConnectionRequestBodyContentStyle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateVpnConnectionRequestBodyContentHaRole struct {
	value string
}

type CreateVpnConnectionRequestBodyContentHaRoleEnum struct {
	MASTER CreateVpnConnectionRequestBodyContentHaRole
	SLAVE  CreateVpnConnectionRequestBodyContentHaRole
}

func GetCreateVpnConnectionRequestBodyContentHaRoleEnum() CreateVpnConnectionRequestBodyContentHaRoleEnum {
	return CreateVpnConnectionRequestBodyContentHaRoleEnum{
		MASTER: CreateVpnConnectionRequestBodyContentHaRole{
			value: "master",
		},
		SLAVE: CreateVpnConnectionRequestBodyContentHaRole{
			value: "slave",
		},
	}
}

func (c CreateVpnConnectionRequestBodyContentHaRole) Value() string {
	return c.value
}

func (c CreateVpnConnectionRequestBodyContentHaRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpnConnectionRequestBodyContentHaRole) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
