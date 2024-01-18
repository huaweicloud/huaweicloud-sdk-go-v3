package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResponseVpnConnection struct {

	// VPN连接ID
	Id *string `json:"id,omitempty"`

	// VPN连接名称
	Name *string `json:"name,omitempty"`

	// VPN连接状态
	Status *string `json:"status,omitempty"`

	// VPN网关ID
	VgwId *string `json:"vgw_id,omitempty"`

	// VGW IP
	VgwIp *string `json:"vgw_ip,omitempty"`

	// 连接模式 允许范围[POLICY, STATIC, BGP] POLICY: 策略模式 STATIC: 静态路由模式 BGP: bgp路由模式
	Style *ResponseVpnConnectionStyle `json:"style,omitempty"`

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

	// 连接监控ID
	ConnectionMonitorId *string `json:"connection_monitor_id,omitempty"`

	// 连接的HA角色
	HaRole *string `json:"ha_role,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o ResponseVpnConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseVpnConnection struct{}"
	}

	return strings.Join([]string{"ResponseVpnConnection", string(data)}, " ")
}

type ResponseVpnConnectionStyle struct {
	value string
}

type ResponseVpnConnectionStyleEnum struct {
	POLICY ResponseVpnConnectionStyle
	STATIC ResponseVpnConnectionStyle
	BGP    ResponseVpnConnectionStyle
}

func GetResponseVpnConnectionStyleEnum() ResponseVpnConnectionStyleEnum {
	return ResponseVpnConnectionStyleEnum{
		POLICY: ResponseVpnConnectionStyle{
			value: "POLICY",
		},
		STATIC: ResponseVpnConnectionStyle{
			value: "STATIC",
		},
		BGP: ResponseVpnConnectionStyle{
			value: "BGP",
		},
	}
}

func (c ResponseVpnConnectionStyle) Value() string {
	return c.value
}

func (c ResponseVpnConnectionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseVpnConnectionStyle) UnmarshalJSON(b []byte) error {
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
