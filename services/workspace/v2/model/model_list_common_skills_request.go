package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCommonSkillsRequest Request Object
type ListCommonSkillsRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListCommonSkillsRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`

	// 技能分类。
	Category *string `json:"category,omitempty"`

	// 技能状态。
	Status *string `json:"status,omitempty"`

	// 技能名称（模糊匹配slug、display_name和alias_name）。
	SkillName *string `json:"skill_name,omitempty"`
}

func (o ListCommonSkillsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillsRequest struct{}"
	}

	return strings.Join([]string{"ListCommonSkillsRequest", string(data)}, " ")
}

type ListCommonSkillsRequestXLanguage struct {
	value string
}

type ListCommonSkillsRequestXLanguageEnum struct {
	EN_US ListCommonSkillsRequestXLanguage
	ZH_CN ListCommonSkillsRequestXLanguage
}

func GetListCommonSkillsRequestXLanguageEnum() ListCommonSkillsRequestXLanguageEnum {
	return ListCommonSkillsRequestXLanguageEnum{
		EN_US: ListCommonSkillsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListCommonSkillsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListCommonSkillsRequestXLanguage) Value() string {
	return c.value
}

func (c ListCommonSkillsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCommonSkillsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
