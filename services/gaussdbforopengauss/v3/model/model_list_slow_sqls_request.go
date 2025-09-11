package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSlowSqlsRequest Request Object
type ListSlowSqlsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSlowSqlsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListSlowSqlsRequestBody `json:"body,omitempty"`
}

func (o ListSlowSqlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowSqlsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowSqlsRequest", string(data)}, " ")
}

type ListSlowSqlsRequestXLanguage struct {
	value string
}

type ListSlowSqlsRequestXLanguageEnum struct {
	ZH_CN ListSlowSqlsRequestXLanguage
	EN_US ListSlowSqlsRequestXLanguage
}

func GetListSlowSqlsRequestXLanguageEnum() ListSlowSqlsRequestXLanguageEnum {
	return ListSlowSqlsRequestXLanguageEnum{
		ZH_CN: ListSlowSqlsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSlowSqlsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSlowSqlsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSlowSqlsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowSqlsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
