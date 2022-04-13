package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 操作
type PolicyAction struct {
	// 防护等级（log为仅记录、block为拦截）

	Category *PolicyActionCategory `json:"category,omitempty"`
}

func (o PolicyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAction struct{}"
	}

	return strings.Join([]string{"PolicyAction", string(data)}, " ")
}

type PolicyActionCategory struct {
	value string
}

type PolicyActionCategoryEnum struct {
	BLOCK PolicyActionCategory
	LOG   PolicyActionCategory
}

func GetPolicyActionCategoryEnum() PolicyActionCategoryEnum {
	return PolicyActionCategoryEnum{
		BLOCK: PolicyActionCategory{
			value: "block",
		},
		LOG: PolicyActionCategory{
			value: "log",
		},
	}
}

func (c PolicyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyActionCategory) UnmarshalJSON(b []byte) error {
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
