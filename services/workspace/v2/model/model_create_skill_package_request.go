package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSkillPackageRequest Request Object
type CreateSkillPackageRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *CreateSkillPackageRequestXLanguage `json:"X-Language,omitempty"`

	// 幂等性标识，UUID格式。 创建类接口携带该请求头，服务端据此实现幂等控制；响应头返回相同值。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	Body *CreateSkillPackageReq `json:"body,omitempty"`
}

func (o CreateSkillPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillPackageRequest struct{}"
	}

	return strings.Join([]string{"CreateSkillPackageRequest", string(data)}, " ")
}

type CreateSkillPackageRequestXLanguage struct {
	value string
}

type CreateSkillPackageRequestXLanguageEnum struct {
	EN_US CreateSkillPackageRequestXLanguage
	ZH_CN CreateSkillPackageRequestXLanguage
}

func GetCreateSkillPackageRequestXLanguageEnum() CreateSkillPackageRequestXLanguageEnum {
	return CreateSkillPackageRequestXLanguageEnum{
		EN_US: CreateSkillPackageRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateSkillPackageRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateSkillPackageRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSkillPackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSkillPackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
