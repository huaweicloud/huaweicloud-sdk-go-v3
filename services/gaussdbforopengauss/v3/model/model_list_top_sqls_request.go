package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopSqlsRequest Request Object
type ListTopSqlsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListTopSqlsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListTopSqlsRequestBody `json:"body,omitempty"`
}

func (o ListTopSqlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSqlsRequest struct{}"
	}

	return strings.Join([]string{"ListTopSqlsRequest", string(data)}, " ")
}

type ListTopSqlsRequestXLanguage struct {
	value string
}

type ListTopSqlsRequestXLanguageEnum struct {
	ZH_CN ListTopSqlsRequestXLanguage
	EN_US ListTopSqlsRequestXLanguage
}

func GetListTopSqlsRequestXLanguageEnum() ListTopSqlsRequestXLanguageEnum {
	return ListTopSqlsRequestXLanguageEnum{
		ZH_CN: ListTopSqlsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTopSqlsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTopSqlsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTopSqlsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopSqlsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
