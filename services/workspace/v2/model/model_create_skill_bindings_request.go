package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSkillBindingsRequest Request Object
type CreateSkillBindingsRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *CreateSkillBindingsRequestXLanguage `json:"X-Language,omitempty"`

	// 幂等性标识，UUID格式。 创建类接口携带该请求头，服务端据此实现幂等控制；响应头返回相同值。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateSkillBindingReq `json:"body,omitempty"`
}

func (o CreateSkillBindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillBindingsRequest struct{}"
	}

	return strings.Join([]string{"CreateSkillBindingsRequest", string(data)}, " ")
}

type CreateSkillBindingsRequestXLanguage struct {
	value string
}

type CreateSkillBindingsRequestXLanguageEnum struct {
	EN_US CreateSkillBindingsRequestXLanguage
	ZH_CN CreateSkillBindingsRequestXLanguage
}

func GetCreateSkillBindingsRequestXLanguageEnum() CreateSkillBindingsRequestXLanguageEnum {
	return CreateSkillBindingsRequestXLanguageEnum{
		EN_US: CreateSkillBindingsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateSkillBindingsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateSkillBindingsRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSkillBindingsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSkillBindingsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
