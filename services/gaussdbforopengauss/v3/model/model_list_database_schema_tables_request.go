package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatabaseSchemaTablesRequest Request Object
type ListDatabaseSchemaTablesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListDatabaseSchemaTablesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 数据库名称。 **约束限制**: 不能使用模板库，且是已存在的数据库名称。 模板库包括postgres， template0，template1，templatea，template_pdb，templatem。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DbName string `json:"db_name"`

	// **参数解释**: SCHEMA名称。 **约束限制**: 不能使用public，information_schema，且schema名称必须存在。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SchemaName string `json:"schema_name"`

	// **参数解释**: 表名关键字。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TableNameKeyword *string `json:"table_name_keyword,omitempty"`

	// **参数解释**: 偏移量，表示从此偏移量开始查询。 **约束限制**: 不涉及。 **取值范围**: [0~2147483647) **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示的条目数量。 **约束限制**: 不涉及。 **取值范围**: [1, 200] **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseSchemaTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseSchemaTablesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseSchemaTablesRequest", string(data)}, " ")
}

type ListDatabaseSchemaTablesRequestXLanguage struct {
	value string
}

type ListDatabaseSchemaTablesRequestXLanguageEnum struct {
	ZH_CN ListDatabaseSchemaTablesRequestXLanguage
	EN_US ListDatabaseSchemaTablesRequestXLanguage
}

func GetListDatabaseSchemaTablesRequestXLanguageEnum() ListDatabaseSchemaTablesRequestXLanguageEnum {
	return ListDatabaseSchemaTablesRequestXLanguageEnum{
		ZH_CN: ListDatabaseSchemaTablesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDatabaseSchemaTablesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDatabaseSchemaTablesRequestXLanguage) Value() string {
	return c.value
}

func (c ListDatabaseSchemaTablesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseSchemaTablesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
