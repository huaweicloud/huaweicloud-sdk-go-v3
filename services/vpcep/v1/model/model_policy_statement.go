package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyStatement policy
type PolicyStatement struct {

	// - Allow，允许控制访问权限 - Deny，拒绝控制访问权限
	Effect PolicyStatementEffect `json:"Effect"`

	// obs访问权限
	Action []string `json:"Action"`

	// obs对象
	Resource []string `json:"Resource"`
}

func (o PolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatement struct{}"
	}

	return strings.Join([]string{"PolicyStatement", string(data)}, " ")
}

type PolicyStatementEffect struct {
	value string
}

type PolicyStatementEffectEnum struct {
	ALLOW PolicyStatementEffect
	DENY  PolicyStatementEffect
}

func GetPolicyStatementEffectEnum() PolicyStatementEffectEnum {
	return PolicyStatementEffectEnum{
		ALLOW: PolicyStatementEffect{
			value: "Allow",
		},
		DENY: PolicyStatementEffect{
			value: "Deny",
		},
	}
}

func (c PolicyStatementEffect) Value() string {
	return c.value
}

func (c PolicyStatementEffect) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyStatementEffect) UnmarshalJSON(b []byte) error {
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
