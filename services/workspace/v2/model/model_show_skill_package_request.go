package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSkillPackageRequest Request Object
type ShowSkillPackageRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ShowSkillPackageRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`
}

func (o ShowSkillPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillPackageRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillPackageRequest", string(data)}, " ")
}

type ShowSkillPackageRequestXLanguage struct {
	value string
}

type ShowSkillPackageRequestXLanguageEnum struct {
	EN_US ShowSkillPackageRequestXLanguage
	ZH_CN ShowSkillPackageRequestXLanguage
}

func GetShowSkillPackageRequestXLanguageEnum() ShowSkillPackageRequestXLanguageEnum {
	return ShowSkillPackageRequestXLanguageEnum{
		EN_US: ShowSkillPackageRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSkillPackageRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSkillPackageRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSkillPackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSkillPackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
