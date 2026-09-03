package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExpandSkillPackageRegionRequest Request Object
type ExpandSkillPackageRegionRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ExpandSkillPackageRegionRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`

	Body *ExpandSkillPackageRegionReq `json:"body,omitempty"`
}

func (o ExpandSkillPackageRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandSkillPackageRegionRequest struct{}"
	}

	return strings.Join([]string{"ExpandSkillPackageRegionRequest", string(data)}, " ")
}

type ExpandSkillPackageRegionRequestXLanguage struct {
	value string
}

type ExpandSkillPackageRegionRequestXLanguageEnum struct {
	EN_US ExpandSkillPackageRegionRequestXLanguage
	ZH_CN ExpandSkillPackageRegionRequestXLanguage
}

func GetExpandSkillPackageRegionRequestXLanguageEnum() ExpandSkillPackageRegionRequestXLanguageEnum {
	return ExpandSkillPackageRegionRequestXLanguageEnum{
		EN_US: ExpandSkillPackageRegionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExpandSkillPackageRegionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExpandSkillPackageRegionRequestXLanguage) Value() string {
	return c.value
}

func (c ExpandSkillPackageRegionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExpandSkillPackageRegionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
