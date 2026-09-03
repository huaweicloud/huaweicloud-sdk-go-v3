package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSkillPackageRequest Request Object
type DeleteSkillPackageRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *DeleteSkillPackageRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`
}

func (o DeleteSkillPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillPackageRequest struct{}"
	}

	return strings.Join([]string{"DeleteSkillPackageRequest", string(data)}, " ")
}

type DeleteSkillPackageRequestXLanguage struct {
	value string
}

type DeleteSkillPackageRequestXLanguageEnum struct {
	EN_US DeleteSkillPackageRequestXLanguage
	ZH_CN DeleteSkillPackageRequestXLanguage
}

func GetDeleteSkillPackageRequestXLanguageEnum() DeleteSkillPackageRequestXLanguageEnum {
	return DeleteSkillPackageRequestXLanguageEnum{
		EN_US: DeleteSkillPackageRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteSkillPackageRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteSkillPackageRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSkillPackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSkillPackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
