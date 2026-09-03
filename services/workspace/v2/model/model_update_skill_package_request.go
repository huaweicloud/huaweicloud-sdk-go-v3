package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSkillPackageRequest Request Object
type UpdateSkillPackageRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *UpdateSkillPackageRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`

	Body *UpdateSkillPackageReq `json:"body,omitempty"`
}

func (o UpdateSkillPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSkillPackageRequest struct{}"
	}

	return strings.Join([]string{"UpdateSkillPackageRequest", string(data)}, " ")
}

type UpdateSkillPackageRequestXLanguage struct {
	value string
}

type UpdateSkillPackageRequestXLanguageEnum struct {
	EN_US UpdateSkillPackageRequestXLanguage
	ZH_CN UpdateSkillPackageRequestXLanguage
}

func GetUpdateSkillPackageRequestXLanguageEnum() UpdateSkillPackageRequestXLanguageEnum {
	return UpdateSkillPackageRequestXLanguageEnum{
		EN_US: UpdateSkillPackageRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateSkillPackageRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateSkillPackageRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSkillPackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSkillPackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
