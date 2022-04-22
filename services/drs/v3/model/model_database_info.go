package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库对象选择信息
type DatabaseInfo struct {

	// object_type为database时，为库名；object_type为table或者view时，字段值参考示例。
	Id string `json:"id"`

	// object_type为table或view时需要填写，为库名
	ParentId *string `json:"parent_id,omitempty"`

	// 类型
	ObjectType DatabaseInfoObjectType `json:"object_type"`

	// 数据库对象名称，库名、表名、视图名
	ObjectName string `json:"object_name"`

	// 别名，映射的新名称。
	ObjectAliasName *string `json:"object_alias_name,omitempty"`

	// 是否选中，值为true会进行迁移，false该数据库对象不会迁移，partial为迁移库下面的部分表，不填默认为false
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

func (c DatabaseInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInfoObjectType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
