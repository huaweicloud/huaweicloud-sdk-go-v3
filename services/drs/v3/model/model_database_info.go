package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatabaseInfo 数据库对象选择信息
type DatabaseInfo struct {

	// **参数解释**：  数据库对象的标识，具体含义取决于object_type的取值。object_type为database时，为库名；object_type为table或view时，字段值参考请求示例。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  数据库对象的父级标识，object_type为table或view时需要填写，为库名。  **约束限制**：  - object_type为table或view时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**：  数据库对象的类型，用于标识迁移或同步对象是数据库、表、模式还是视图。  **约束限制**：  不涉及。  **取值范围**：  - database：数据库 - table：表 - schema：模式 - view：视图  **默认取值**：  不涉及
	ObjectType *DatabaseInfoObjectType `json:"object_type,omitempty"`

	// **参数解释**：  数据库对象名称，可以是库名、表名或视图名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ObjectName *string `json:"object_name,omitempty"`

	// **参数解释**：  数据库对象的别名，即映射后的新名称。  **约束限制**：  - 仅支持同步任务使用。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	// **参数解释**：  是否选中该数据库对象进行迁移，用于控制对象级别的迁移选择。  **约束限制**：  不涉及。  **取值范围**：  - true：选中，该数据库对象会进行迁移 - false：不选中，该数据库对象不会迁移 - partial：部分选择，迁移库下面的部分表  **默认取值**：  false
	Select *string `json:"select,omitempty"`
}

func (o DatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInfo struct{}"
	}

	return strings.Join([]string{"DatabaseInfo", string(data)}, " ")
}

type DatabaseInfoObjectType struct {
	value string
}

type DatabaseInfoObjectTypeEnum struct {
	DATABASE DatabaseInfoObjectType
	TABLE    DatabaseInfoObjectType
	SCHEMA   DatabaseInfoObjectType
	VIEW     DatabaseInfoObjectType
}

func GetDatabaseInfoObjectTypeEnum() DatabaseInfoObjectTypeEnum {
	return DatabaseInfoObjectTypeEnum{
		DATABASE: DatabaseInfoObjectType{
			value: "database",
		},
		TABLE: DatabaseInfoObjectType{
			value: "table",
		},
		SCHEMA: DatabaseInfoObjectType{
			value: "schema",
		},
		VIEW: DatabaseInfoObjectType{
			value: "view",
		},
	}
}

func (c DatabaseInfoObjectType) Value() string {
	return c.value
}

func (c DatabaseInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInfoObjectType) UnmarshalJSON(b []byte) error {
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
