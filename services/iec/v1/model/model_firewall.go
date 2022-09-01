package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 防火墙对象
type Firewall struct {

	// 网络ACL ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 网络ACL名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 网络ACL使能开关。  取值范围：true（开启），false（关闭）。默认为true
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	// 网络ACL状态。  取值范围：INACTIVE
	Status *FirewallStatus `json:"status,omitempty" xml:"status"`

	// 网络ACL描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 租户domainID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	EgressFirewallPolicy *FirewallPolicy `json:"egress_firewall_policy,omitempty" xml:"egress_firewall_policy"`

	// 出方向网络ACL规则个数。
	EgressFirewallRuleCount *int32 `json:"egress_firewall_rule_count,omitempty" xml:"egress_firewall_rule_count"`

	IngressFirewallPolicy *FirewallPolicy `json:"ingress_firewall_policy,omitempty" xml:"ingress_firewall_policy"`

	// 入方向网络ACL规则个数。
	IngressFirewallRuleCount *int32 `json:"ingress_firewall_rule_count,omitempty" xml:"ingress_firewall_rule_count"`
}

func (o Firewall) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Firewall struct{}"
	}

	return strings.Join([]string{"Firewall", string(data)}, " ")
}

type FirewallStatus struct {
	value string
}

type FirewallStatusEnum struct {
	INACTIVE FirewallStatus
}

func GetFirewallStatusEnum() FirewallStatusEnum {
	return FirewallStatusEnum{
		INACTIVE: FirewallStatus{
			value: "INACTIVE",
		},
	}
}

func (c FirewallStatus) Value() string {
	return c.value
}

func (c FirewallStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FirewallStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
