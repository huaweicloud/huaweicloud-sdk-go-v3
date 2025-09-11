package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableDefinitionsRequest Request Object
type ListTableDefinitionsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListTableDefinitionsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释** : 数据库名称。 **约束限制** : 不涉及。 **取值范围** : 只能由英文字母、数字组成，且长度为36个字符。 **默认取值** : 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释** : schema名称。 **约束限制** : 不涉及。 **取值范围** : 只能由英文字母、数字组成，且长度为36个字符。 **默认取值** : 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释** : 表名称。 **约束限制** : 不涉及。 **取值范围** : 只能由英文字母、数字组成，且长度为36个字符。 **默认取值** : 不涉及。
	TableName string `json:"table_name"`
}

func (o ListTableDefinitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableDefinitionsRequest struct{}"
	}

	return strings.Join([]string{"ListTableDefinitionsRequest", string(data)}, " ")
}

type ListTableDefinitionsRequestXLanguage struct {
	value string
}

type ListTableDefinitionsRequestXLanguageEnum struct {
	ZH_CN ListTableDefinitionsRequestXLanguage
	EN_US ListTableDefinitionsRequestXLanguage
}

func GetListTableDefinitionsRequestXLanguageEnum() ListTableDefinitionsRequestXLanguageEnum {
	return ListTableDefinitionsRequestXLanguageEnum{
		ZH_CN: ListTableDefinitionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTableDefinitionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTableDefinitionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTableDefinitionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableDefinitionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
