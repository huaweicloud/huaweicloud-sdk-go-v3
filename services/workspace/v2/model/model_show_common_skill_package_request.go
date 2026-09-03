package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCommonSkillPackageRequest Request Object
type ShowCommonSkillPackageRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ShowCommonSkillPackageRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`
}

func (o ShowCommonSkillPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommonSkillPackageRequest struct{}"
	}

	return strings.Join([]string{"ShowCommonSkillPackageRequest", string(data)}, " ")
}

type ShowCommonSkillPackageRequestXLanguage struct {
	value string
}

type ShowCommonSkillPackageRequestXLanguageEnum struct {
	EN_US ShowCommonSkillPackageRequestXLanguage
	ZH_CN ShowCommonSkillPackageRequestXLanguage
}

func GetShowCommonSkillPackageRequestXLanguageEnum() ShowCommonSkillPackageRequestXLanguageEnum {
	return ShowCommonSkillPackageRequestXLanguageEnum{
		EN_US: ShowCommonSkillPackageRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowCommonSkillPackageRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowCommonSkillPackageRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCommonSkillPackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCommonSkillPackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
