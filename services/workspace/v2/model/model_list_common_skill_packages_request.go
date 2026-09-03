package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCommonSkillPackagesRequest Request Object
type ListCommonSkillPackagesRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListCommonSkillPackagesRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCommonSkillPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillPackagesRequest struct{}"
	}

	return strings.Join([]string{"ListCommonSkillPackagesRequest", string(data)}, " ")
}

type ListCommonSkillPackagesRequestXLanguage struct {
	value string
}

type ListCommonSkillPackagesRequestXLanguageEnum struct {
	EN_US ListCommonSkillPackagesRequestXLanguage
	ZH_CN ListCommonSkillPackagesRequestXLanguage
}

func GetListCommonSkillPackagesRequestXLanguageEnum() ListCommonSkillPackagesRequestXLanguageEnum {
	return ListCommonSkillPackagesRequestXLanguageEnum{
		EN_US: ListCommonSkillPackagesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListCommonSkillPackagesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListCommonSkillPackagesRequestXLanguage) Value() string {
	return c.value
}

func (c ListCommonSkillPackagesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCommonSkillPackagesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
