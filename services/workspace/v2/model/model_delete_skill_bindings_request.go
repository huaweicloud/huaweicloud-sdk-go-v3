package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSkillBindingsRequest Request Object
type DeleteSkillBindingsRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *DeleteSkillBindingsRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteSkillBindingReq `json:"body,omitempty"`
}

func (o DeleteSkillBindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillBindingsRequest struct{}"
	}

	return strings.Join([]string{"DeleteSkillBindingsRequest", string(data)}, " ")
}

type DeleteSkillBindingsRequestXLanguage struct {
	value string
}

type DeleteSkillBindingsRequestXLanguageEnum struct {
	EN_US DeleteSkillBindingsRequestXLanguage
	ZH_CN DeleteSkillBindingsRequestXLanguage
}

func GetDeleteSkillBindingsRequestXLanguageEnum() DeleteSkillBindingsRequestXLanguageEnum {
	return DeleteSkillBindingsRequestXLanguageEnum{
		EN_US: DeleteSkillBindingsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteSkillBindingsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteSkillBindingsRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSkillBindingsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSkillBindingsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
