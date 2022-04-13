package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 预览邮件格式请求体
type PreviewTemplateBody struct {
	// 邮件模板内容

	Templates string `json:"templates"`
	// 语言 zh-cn中文，en-us英文

	Language PreviewTemplateBodyLanguage `json:"language"`
	// 来源，只能填LTS

	Source string `json:"source"`
}

func (o PreviewTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewTemplateBody struct{}"
	}

	return strings.Join([]string{"PreviewTemplateBody", string(data)}, " ")
}

type PreviewTemplateBodyLanguage struct {
	value string
}

type PreviewTemplateBodyLanguageEnum struct {
	ZH_CN PreviewTemplateBodyLanguage
	EN_US PreviewTemplateBodyLanguage
}

func GetPreviewTemplateBodyLanguageEnum() PreviewTemplateBodyLanguageEnum {
	return PreviewTemplateBodyLanguageEnum{
		ZH_CN: PreviewTemplateBodyLanguage{
			value: "zh-cn",
		},
		EN_US: PreviewTemplateBodyLanguage{
			value: "en-us",
		},
	}
}

func (c PreviewTemplateBodyLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewTemplateBodyLanguage) UnmarshalJSON(b []byte) error {
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
