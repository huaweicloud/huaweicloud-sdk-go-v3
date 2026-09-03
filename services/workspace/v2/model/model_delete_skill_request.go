package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSkillRequest Request Object
type DeleteSkillRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *DeleteSkillRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`
}

func (o DeleteSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillRequest struct{}"
	}

	return strings.Join([]string{"DeleteSkillRequest", string(data)}, " ")
}

type DeleteSkillRequestXLanguage struct {
	value string
}

type DeleteSkillRequestXLanguageEnum struct {
	EN_US DeleteSkillRequestXLanguage
	ZH_CN DeleteSkillRequestXLanguage
}

func GetDeleteSkillRequestXLanguageEnum() DeleteSkillRequestXLanguageEnum {
	return DeleteSkillRequestXLanguageEnum{
		EN_US: DeleteSkillRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteSkillRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteSkillRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSkillRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSkillRequestXLanguage) UnmarshalJSON(b []byte) error {
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
