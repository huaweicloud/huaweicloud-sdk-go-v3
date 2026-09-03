package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCommonSkillResourcesRequest Request Object
type ListCommonSkillResourcesRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListCommonSkillResourcesRequestXLanguage `json:"X-Language,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能版本号。
	Version *string `json:"version,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCommonSkillResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListCommonSkillResourcesRequest", string(data)}, " ")
}

type ListCommonSkillResourcesRequestXLanguage struct {
	value string
}

type ListCommonSkillResourcesRequestXLanguageEnum struct {
	EN_US ListCommonSkillResourcesRequestXLanguage
	ZH_CN ListCommonSkillResourcesRequestXLanguage
}

func GetListCommonSkillResourcesRequestXLanguageEnum() ListCommonSkillResourcesRequestXLanguageEnum {
	return ListCommonSkillResourcesRequestXLanguageEnum{
		EN_US: ListCommonSkillResourcesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListCommonSkillResourcesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListCommonSkillResourcesRequestXLanguage) Value() string {
	return c.value
}

func (c ListCommonSkillResourcesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCommonSkillResourcesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
