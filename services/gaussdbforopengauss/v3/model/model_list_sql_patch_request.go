package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlPatchRequest Request Object
type ListSqlPatchRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *ListSqlPatchRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListSqlPatchRequestBody `json:"body,omitempty"`
}

func (o ListSqlPatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlPatchRequest struct{}"
	}

	return strings.Join([]string{"ListSqlPatchRequest", string(data)}, " ")
}

type ListSqlPatchRequestXLanguage struct {
	value string
}

type ListSqlPatchRequestXLanguageEnum struct {
	ZH_CN ListSqlPatchRequestXLanguage
	EN_US ListSqlPatchRequestXLanguage
}

func GetListSqlPatchRequestXLanguageEnum() ListSqlPatchRequestXLanguageEnum {
	return ListSqlPatchRequestXLanguageEnum{
		ZH_CN: ListSqlPatchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlPatchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlPatchRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlPatchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlPatchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
