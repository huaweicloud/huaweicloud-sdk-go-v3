package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableDefinitionRequest Request Object
type ListTableDefinitionRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListTableDefinitionRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 数据库名称。
	DatabaseName string `json:"database_name"`

	// schema名称。
	SchemaName string `json:"schema_name"`

	// 表名称。
	TableName string `json:"table_name"`
}

func (o ListTableDefinitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableDefinitionRequest struct{}"
	}

	return strings.Join([]string{"ListTableDefinitionRequest", string(data)}, " ")
}

type ListTableDefinitionRequestXLanguage struct {
	value string
}

type ListTableDefinitionRequestXLanguageEnum struct {
	ZH_CN ListTableDefinitionRequestXLanguage
	EN_US ListTableDefinitionRequestXLanguage
}

func GetListTableDefinitionRequestXLanguageEnum() ListTableDefinitionRequestXLanguageEnum {
	return ListTableDefinitionRequestXLanguageEnum{
		ZH_CN: ListTableDefinitionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTableDefinitionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTableDefinitionRequestXLanguage) Value() string {
	return c.value
}

func (c ListTableDefinitionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableDefinitionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
