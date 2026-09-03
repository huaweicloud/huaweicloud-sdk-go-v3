package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSkillPackagesRequest Request Object
type ListSkillPackagesRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListSkillPackagesRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSkillPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSkillPackagesRequest struct{}"
	}

	return strings.Join([]string{"ListSkillPackagesRequest", string(data)}, " ")
}

type ListSkillPackagesRequestXLanguage struct {
	value string
}

type ListSkillPackagesRequestXLanguageEnum struct {
	EN_US ListSkillPackagesRequestXLanguage
	ZH_CN ListSkillPackagesRequestXLanguage
}

func GetListSkillPackagesRequestXLanguageEnum() ListSkillPackagesRequestXLanguageEnum {
	return ListSkillPackagesRequestXLanguageEnum{
		EN_US: ListSkillPackagesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListSkillPackagesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListSkillPackagesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSkillPackagesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSkillPackagesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
