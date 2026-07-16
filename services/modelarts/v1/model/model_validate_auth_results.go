package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ValidateAuthResults struct {

	// **参数解释**：细粒度权限。 **取值范围**：不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**：鉴权通过与否。 **取值范围**： - allow：通过。 - deny：不通过。
	Verdict *ValidateAuthResultsVerdict `json:"verdict,omitempty"`

	// **参数解释**：随机的uuid，用来定位问题使用。 **取值范围**：不涉及。
	ActionId *string `json:"action_id,omitempty"`

	// **参数解释**：请求资源。 **取值范围**：不涉及。
	Resource *string `json:"resource,omitempty"`

	// **参数解释**：失败情况下原因。
	Cause *[]Cause `json:"cause,omitempty"`
}

func (o ValidateAuthResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAuthResults struct{}"
	}

	return strings.Join([]string{"ValidateAuthResults", string(data)}, " ")
}

type ValidateAuthResultsVerdict struct {
	value string
}

type ValidateAuthResultsVerdictEnum struct {
	ALLOW ValidateAuthResultsVerdict
	DENY  ValidateAuthResultsVerdict
}

func GetValidateAuthResultsVerdictEnum() ValidateAuthResultsVerdictEnum {
	return ValidateAuthResultsVerdictEnum{
		ALLOW: ValidateAuthResultsVerdict{
			value: "allow",
		},
		DENY: ValidateAuthResultsVerdict{
			value: "deny",
		},
	}
}

func (c ValidateAuthResultsVerdict) Value() string {
	return c.value
}

func (c ValidateAuthResultsVerdict) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateAuthResultsVerdict) UnmarshalJSON(b []byte) error {
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
