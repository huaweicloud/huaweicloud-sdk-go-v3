package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateFirewallRuleResp
type UpdateFirewallRuleResp struct {

	// 网络ACL ID
	Id string `json:"id"`

	// 网络ACL状态。
	Status UpdateFirewallRuleRespStatus `json:"status"`

	EgressFirewallPolicy *FirewallPolicy `json:"egress_firewall_policy"`

	IngressFirewallPolicy *FirewallPolicy `json:"ingress_firewall_policy"`
}

func (o UpdateFirewallRuleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleResp struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleResp", string(data)}, " ")
}

type UpdateFirewallRuleRespStatus struct {
	value string
}

type UpdateFirewallRuleRespStatusEnum struct {
	INACTIVE UpdateFirewallRuleRespStatus
}

func GetUpdateFirewallRuleRespStatusEnum() UpdateFirewallRuleRespStatusEnum {
	return UpdateFirewallRuleRespStatusEnum{
		INACTIVE: UpdateFirewallRuleRespStatus{
			value: "INACTIVE",
		},
	}
}

func (c UpdateFirewallRuleRespStatus) Value() string {
	return c.value
}

func (c UpdateFirewallRuleRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFirewallRuleRespStatus) UnmarshalJSON(b []byte) error {
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
