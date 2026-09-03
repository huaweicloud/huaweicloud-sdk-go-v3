package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSkillRequest Request Object
type UpdateSkillRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *UpdateSkillRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	Body *UpdateSkillReq `json:"body,omitempty"`
}

func (o UpdateSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSkillRequest struct{}"
	}

	return strings.Join([]string{"UpdateSkillRequest", string(data)}, " ")
}

type UpdateSkillRequestXLanguage struct {
	value string
}

type UpdateSkillRequestXLanguageEnum struct {
	EN_US UpdateSkillRequestXLanguage
	ZH_CN UpdateSkillRequestXLanguage
}

func GetUpdateSkillRequestXLanguageEnum() UpdateSkillRequestXLanguageEnum {
	return UpdateSkillRequestXLanguageEnum{
		EN_US: UpdateSkillRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateSkillRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateSkillRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSkillRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSkillRequestXLanguage) UnmarshalJSON(b []byte) error {
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
