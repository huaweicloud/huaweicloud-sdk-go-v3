package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSkillRequest Request Object
type CreateSkillRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *CreateSkillRequestXLanguage `json:"X-Language,omitempty"`

	// 幂等性标识，UUID格式。 创建类接口携带该请求头，服务端据此实现幂等控制；响应头返回相同值。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateSkillReq `json:"body,omitempty"`
}

func (o CreateSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillRequest struct{}"
	}

	return strings.Join([]string{"CreateSkillRequest", string(data)}, " ")
}

type CreateSkillRequestXLanguage struct {
	value string
}

type CreateSkillRequestXLanguageEnum struct {
	EN_US CreateSkillRequestXLanguage
	ZH_CN CreateSkillRequestXLanguage
}

func GetCreateSkillRequestXLanguageEnum() CreateSkillRequestXLanguageEnum {
	return CreateSkillRequestXLanguageEnum{
		EN_US: CreateSkillRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateSkillRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateSkillRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSkillRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSkillRequestXLanguage) UnmarshalJSON(b []byte) error {
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
