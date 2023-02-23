package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库信息
type DatabaseInput struct {

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 数据库所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型
	OwnerType *DatabaseInputOwnerType `json:"owner_type,omitempty"`

	// 数据库描述信息
	Description *string `json:"description,omitempty"`

	// 数据库位置
	Location *string `json:"location,omitempty"`

	// 标签信息
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表路径列表
	TableLocationList *[]string `json:"table_location_list,omitempty"`

	// 函数路径列表
	FunctionLocationList *[]string `json:"function_location_list,omitempty"`
}

func (o DatabaseInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInput struct{}"
	}

	return strings.Join([]string{"DatabaseInput", string(data)}, " ")
}

type DatabaseInputOwnerType struct {
	value string
}

type DatabaseInputOwnerTypeEnum struct {
	USER  DatabaseInputOwnerType
	ROLE  DatabaseInputOwnerType
	GROUP DatabaseInputOwnerType
}

func GetDatabaseInputOwnerTypeEnum() DatabaseInputOwnerTypeEnum {
	return DatabaseInputOwnerTypeEnum{
		USER: DatabaseInputOwnerType{
			value: "USER",
		},
		ROLE: DatabaseInputOwnerType{
			value: "ROLE",
		},
		GROUP: DatabaseInputOwnerType{
			value: "GROUP",
		},
	}
}

func (c DatabaseInputOwnerType) Value() string {
	return c.value
}

func (c DatabaseInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInputOwnerType) UnmarshalJSON(b []byte) error {
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
