package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"bytes"
	"encoding/json"
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

	// statement生效的条件
	Condition *interface{} `json:"Condition,omitempty"`

	// 策略Id
	Sid *string `json:"Sid,omitempty"`
}

func (o PolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatement struct{}"
	}

	return strings.Join([]string{"PolicyStatement", string(data)}, " ")
}

// UnmarshalJSON accepts both the single-string and the array-of-strings forms
// of Action and Resource: gateway endpoints (OBS/SFS) return their default
// full-access policy with a plain string, e.g. {"Effect": "Allow", "Action": "*"}.
func (m *PolicyStatement) UnmarshalJSON(data []byte) error {
	type PolicyStatementAlias PolicyStatement
	aux := &struct {
		Action   json.RawMessage `json:"Action"`
		Resource json.RawMessage `json:"Resource"`
		*PolicyStatementAlias
	}{
		PolicyStatementAlias: (*PolicyStatementAlias)(m),
	}
	if err := utils.Unmarshal(data, aux); err != nil {
		return err
	}

	action, err := decodeStringOrStringSlice(aux.Action)
	if err != nil {
		return err
	}
	resource, err := decodeStringOrStringSlice(aux.Resource)
	if err != nil {
		return err
	}
	m.Action = action
	m.Resource = resource
	return nil
}

func decodeStringOrStringSlice(data []byte) ([]string, error) {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil, nil
	}
	if trimmed[0] == '"' {
		var single string
		if err := json.Unmarshal(trimmed, &single); err != nil {
			return nil, err
		}
		return []string{single}, nil
	}

	var slice []string
	if err := json.Unmarshal(trimmed, &slice); err != nil {
		return nil, err
	}
	return slice, nil
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
