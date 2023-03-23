package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 精准防护规则命中后操作对象
type CustomAction struct {

	// 操作类型。   - “block”：拦截。   - “pass”：放行。   - “log”：仅记录
	Category CustomActionCategory `json:"category"`

	// 攻击惩罚规则id，只有当category参数值为block时才可配置该参数
	FollowedActionId *string `json:"followed_action_id,omitempty"`
}

func (o CustomAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomAction struct{}"
	}

	return strings.Join([]string{"CustomAction", string(data)}, " ")
}

type CustomActionCategory struct {
	value string
}

type CustomActionCategoryEnum struct {
	BLOCK CustomActionCategory
	PASS  CustomActionCategory
	LOG   CustomActionCategory
}

func GetCustomActionCategoryEnum() CustomActionCategoryEnum {
	return CustomActionCategoryEnum{
		BLOCK: CustomActionCategory{
			value: "block",
		},
		PASS: CustomActionCategory{
			value: "pass",
		},
		LOG: CustomActionCategory{
			value: "log",
		},
	}
}

func (c CustomActionCategory) Value() string {
	return c.value
}

func (c CustomActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomActionCategory) UnmarshalJSON(b []byte) error {
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
