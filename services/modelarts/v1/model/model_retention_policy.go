package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RetentionPolicy struct {

	// **参数解释**：作业级库记录自动老化策略。 **约束限制**：   - 仅当平台开启作业老化能力且作业类型为自定义训练作业（kind=job）时生效；   - 与用户级「作业自动老化」开关联动：     - 用户级开关**开启**：该用户下所有作业均参与老化（`policy=disabled` 不能单独豁免）；     - 用户级开关**关闭**：仅 `policy=enabled` 的作业参与老化；未设置或 `disabled` 均不参与。 **取值范围**：   - enabled：开启本作业老化   - disabled：关闭本作业老化（仅在用户级开关关闭时有效） **默认取值**：不传表示未单独设置，跟随用户级开关策略。
	Policy *RetentionPolicyPolicy `json:"policy,omitempty"`
}

func (o RetentionPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetentionPolicy struct{}"
	}

	return strings.Join([]string{"RetentionPolicy", string(data)}, " ")
}

type RetentionPolicyPolicy struct {
	value string
}

type RetentionPolicyPolicyEnum struct {
	ENABLED  RetentionPolicyPolicy
	DISABLED RetentionPolicyPolicy
}

func GetRetentionPolicyPolicyEnum() RetentionPolicyPolicyEnum {
	return RetentionPolicyPolicyEnum{
		ENABLED: RetentionPolicyPolicy{
			value: "enabled",
		},
		DISABLED: RetentionPolicyPolicy{
			value: "disabled",
		},
	}
}

func (c RetentionPolicyPolicy) Value() string {
	return c.value
}

func (c RetentionPolicyPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RetentionPolicyPolicy) UnmarshalJSON(b []byte) error {
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
