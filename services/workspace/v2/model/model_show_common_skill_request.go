package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCommonSkillRequest Request Object
type ShowCommonSkillRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ShowCommonSkillRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`
}

func (o ShowCommonSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommonSkillRequest struct{}"
	}

	return strings.Join([]string{"ShowCommonSkillRequest", string(data)}, " ")
}

type ShowCommonSkillRequestXLanguage struct {
	value string
}

type ShowCommonSkillRequestXLanguageEnum struct {
	EN_US ShowCommonSkillRequestXLanguage
	ZH_CN ShowCommonSkillRequestXLanguage
}

func GetShowCommonSkillRequestXLanguageEnum() ShowCommonSkillRequestXLanguageEnum {
	return ShowCommonSkillRequestXLanguageEnum{
		EN_US: ShowCommonSkillRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowCommonSkillRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowCommonSkillRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCommonSkillRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCommonSkillRequestXLanguage) UnmarshalJSON(b []byte) error {
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
