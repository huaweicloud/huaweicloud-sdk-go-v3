package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RemoveSkillPackageRegionRequest Request Object
type RemoveSkillPackageRegionRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *RemoveSkillPackageRegionRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`

	Body *RemovePackageRegionReq `json:"body,omitempty"`
}

func (o RemoveSkillPackageRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveSkillPackageRegionRequest struct{}"
	}

	return strings.Join([]string{"RemoveSkillPackageRegionRequest", string(data)}, " ")
}

type RemoveSkillPackageRegionRequestXLanguage struct {
	value string
}

type RemoveSkillPackageRegionRequestXLanguageEnum struct {
	EN_US RemoveSkillPackageRegionRequestXLanguage
	ZH_CN RemoveSkillPackageRegionRequestXLanguage
}

func GetRemoveSkillPackageRegionRequestXLanguageEnum() RemoveSkillPackageRegionRequestXLanguageEnum {
	return RemoveSkillPackageRegionRequestXLanguageEnum{
		EN_US: RemoveSkillPackageRegionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: RemoveSkillPackageRegionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c RemoveSkillPackageRegionRequestXLanguage) Value() string {
	return c.value
}

func (c RemoveSkillPackageRegionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemoveSkillPackageRegionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
