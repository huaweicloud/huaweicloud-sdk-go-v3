package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Database database
type Database struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 数据库所有者
	Owner string `json:"owner"`

	// 所有者类型
	OwnerType DatabaseOwnerType `json:"owner_type"`

	// 数据库描述信息
	Description string `json:"description"`

	// 数据库路径地址
	Location string `json:"location"`

	// 参数信息
	Parameters map[string]string `json:"parameters"`

	// 表路径列表
	TableLocationList *[]string `json:"table_location_list,omitempty"`

	// 函数路径列表
	FunctionLocationList *[]string `json:"function_location_list,omitempty"`
}

func (o Database) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Database struct{}"
	}

	return strings.Join([]string{"Database", string(data)}, " ")
}

type DatabaseOwnerType struct {
	value string
}

type DatabaseOwnerTypeEnum struct {
	USER  DatabaseOwnerType
	GROUP DatabaseOwnerType
	ROLE  DatabaseOwnerType
}

func GetDatabaseOwnerTypeEnum() DatabaseOwnerTypeEnum {
	return DatabaseOwnerTypeEnum{
		USER: DatabaseOwnerType{
			value: "USER",
		},
		GROUP: DatabaseOwnerType{
			value: "GROUP",
		},
		ROLE: DatabaseOwnerType{
			value: "ROLE",
		},
	}
}

func (c DatabaseOwnerType) Value() string {
	return c.value
}

func (c DatabaseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseOwnerType) UnmarshalJSON(b []byte) error {
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
