package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CcrulesListInfoAction 请求次数限制到达后采取的防护动作
type CcrulesListInfoAction struct {

	// 动作类型：  - captcha：人机验证，阻断后用户需要输入正确的验证码，恢复正确的访问页面。  -block：阻断。   - log: 仅记录   - dynamic_block: 上一个限速周期内，请求频率超过“限速频率”将被阻断，那么在下一个限速周期内，请求频率超过“放行频率”将被阻断。注：只有当cc防护规则模式为高级模式时才支持设置dynamic_block防护动作。
	Category CcrulesListInfoActionCategory `json:"category"`

	Detail *CcrulesListInfoActionDetail `json:"detail,omitempty"`
}

func (o CcrulesListInfoAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcrulesListInfoAction struct{}"
	}

	return strings.Join([]string{"CcrulesListInfoAction", string(data)}, " ")
}

type CcrulesListInfoActionCategory struct {
	value string
}

type CcrulesListInfoActionCategoryEnum struct {
	CAPTCHA       CcrulesListInfoActionCategory
	BLOCK         CcrulesListInfoActionCategory
	LOG           CcrulesListInfoActionCategory
	DYNAMIC_BLOCK CcrulesListInfoActionCategory
}

func GetCcrulesListInfoActionCategoryEnum() CcrulesListInfoActionCategoryEnum {
	return CcrulesListInfoActionCategoryEnum{
		CAPTCHA: CcrulesListInfoActionCategory{
			value: "captcha",
		},
		BLOCK: CcrulesListInfoActionCategory{
			value: "block",
		},
		LOG: CcrulesListInfoActionCategory{
			value: "log",
		},
		DYNAMIC_BLOCK: CcrulesListInfoActionCategory{
			value: "dynamic_block",
		},
	}
}

func (c CcrulesListInfoActionCategory) Value() string {
	return c.value
}

func (c CcrulesListInfoActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CcrulesListInfoActionCategory) UnmarshalJSON(b []byte) error {
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
