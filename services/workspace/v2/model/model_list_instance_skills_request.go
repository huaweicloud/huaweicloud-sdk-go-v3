package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceSkillsRequest Request Object
type ListInstanceSkillsRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *ListInstanceSkillsRequestXLanguage `json:"X-Language,omitempty"`

	// 实例标识。
	InstanceId string `json:"instance_id"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceSkillsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSkillsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceSkillsRequest", string(data)}, " ")
}

type ListInstanceSkillsRequestXLanguage struct {
	value string
}

type ListInstanceSkillsRequestXLanguageEnum struct {
	EN_US ListInstanceSkillsRequestXLanguage
	ZH_CN ListInstanceSkillsRequestXLanguage
}

func GetListInstanceSkillsRequestXLanguageEnum() ListInstanceSkillsRequestXLanguageEnum {
	return ListInstanceSkillsRequestXLanguageEnum{
		EN_US: ListInstanceSkillsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListInstanceSkillsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListInstanceSkillsRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceSkillsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceSkillsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
