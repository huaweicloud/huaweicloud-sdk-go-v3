package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IpReputationRulesListInfoAction **参数解释：** 请求次数限制到达后采取的防护动作 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type IpReputationRulesListInfoAction struct {

	// **参数解释：** 动作类型：  - pass：放行  - block：阻断。   - log: 仅记录  **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Category IpReputationRulesListInfoActionCategory `json:"category"`
}

func (o IpReputationRulesListInfoAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpReputationRulesListInfoAction struct{}"
	}

	return strings.Join([]string{"IpReputationRulesListInfoAction", string(data)}, " ")
}

type IpReputationRulesListInfoActionCategory struct {
	value string
}

type IpReputationRulesListInfoActionCategoryEnum struct {
	PASS  IpReputationRulesListInfoActionCategory
	BLOCK IpReputationRulesListInfoActionCategory
	LOG   IpReputationRulesListInfoActionCategory
}

func GetIpReputationRulesListInfoActionCategoryEnum() IpReputationRulesListInfoActionCategoryEnum {
	return IpReputationRulesListInfoActionCategoryEnum{
		PASS: IpReputationRulesListInfoActionCategory{
			value: "pass",
		},
		BLOCK: IpReputationRulesListInfoActionCategory{
			value: "block",
		},
		LOG: IpReputationRulesListInfoActionCategory{
			value: "log",
		},
	}
}

func (c IpReputationRulesListInfoActionCategory) Value() string {
	return c.value
}

func (c IpReputationRulesListInfoActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpReputationRulesListInfoActionCategory) UnmarshalJSON(b []byte) error {
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
