package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSkillRequest Request Object
type ShowSkillRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ShowSkillRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`
}

func (o ShowSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillRequest", string(data)}, " ")
}

type ShowSkillRequestXLanguage struct {
	value string
}

type ShowSkillRequestXLanguageEnum struct {
	EN_US ShowSkillRequestXLanguage
	ZH_CN ShowSkillRequestXLanguage
}

func GetShowSkillRequestXLanguageEnum() ShowSkillRequestXLanguageEnum {
	return ShowSkillRequestXLanguageEnum{
		EN_US: ShowSkillRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSkillRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSkillRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSkillRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSkillRequestXLanguage) UnmarshalJSON(b []byte) error {
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
