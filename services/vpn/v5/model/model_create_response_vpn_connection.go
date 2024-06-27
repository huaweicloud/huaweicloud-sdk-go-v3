package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateResponseVpnConnection struct {

	// VPN连接ID
	Id *string `json:"id,omitempty"`

	// VPN连接名称
	Name *string `json:"name,omitempty"`

	// VPN网关ID
	VgwId *string `json:"vgw_id,omitempty"`

	// VGW IP
	VgwIp *string `json:"vgw_ip,omitempty"`

	// 连接模式 允许范围[POLICY, STATIC, BGP] POLICY: 策略模式 STATIC: 静态路由模式 BGP: bgp路由模式
	Style *CreateResponseVpnConnectionStyle `json:"style,omitempty"`

	// 对端网关ID
	CgwId *string `json:"cgw_id,omitempty"`

	// 对端网段
	PeerSubnets *[]string `json:"peer_subnets,omitempty"`

	// 本端隧道口地址
	TunnelLocalAddress *string `json:"tunnel_local_address,omitempty"`

	// 对端隧道口地址
	TunnelPeerAddress *string `json:"tunnel_peer_address,omitempty"`

	// 开启NQA检测
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 策略模式的策略规则组
	PolicyRules *[]PolicyRule `json:"policy_rules,omitempty"`

	Ikepolicy *IkePolicy `json:"ikepolicy,omitempty"`

	Ipsecpolicy *IpsecPolicy `json:"ipsecpolicy,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 连接的HA角色
	HaRole *string `json:"ha_role,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o CreateResponseVpnConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResponseVpnConnection struct{}"
	}

	return strings.Join([]string{"CreateResponseVpnConnection", string(data)}, " ")
}

type CreateResponseVpnConnectionStyle struct {
	value string
}

type CreateResponseVpnConnectionStyleEnum struct {
	POLICY CreateResponseVpnConnectionStyle
	STATIC CreateResponseVpnConnectionStyle
	BGP    CreateResponseVpnConnectionStyle
}

func GetCreateResponseVpnConnectionStyleEnum() CreateResponseVpnConnectionStyleEnum {
	return CreateResponseVpnConnectionStyleEnum{
		POLICY: CreateResponseVpnConnectionStyle{
			value: "POLICY",
		},
		STATIC: CreateResponseVpnConnectionStyle{
			value: "STATIC",
		},
		BGP: CreateResponseVpnConnectionStyle{
			value: "BGP",
		},
	}
}

func (c CreateResponseVpnConnectionStyle) Value() string {
	return c.value
}

func (c CreateResponseVpnConnectionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResponseVpnConnectionStyle) UnmarshalJSON(b []byte) error {
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
