package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDatabaseResponse Response Object
type ShowDatabaseResponse struct {

	// catalog名字
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型
	OwnerType *ShowDatabaseResponseOwnerType `json:"owner_type,omitempty"`

	// 数据库描述信息
	Description *string `json:"description,omitempty"`

	// 数据库路径地址
	Location *string `json:"location,omitempty"`

	// 参数信息
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表路径列表
	TableLocationList *[]string `json:"table_location_list,omitempty"`

	// 函数路径列表
	FunctionLocationList *[]string `json:"function_location_list,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o ShowDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseResponse", string(data)}, " ")
}

type ShowDatabaseResponseOwnerType struct {
	value string
}

type ShowDatabaseResponseOwnerTypeEnum struct {
	USER  ShowDatabaseResponseOwnerType
	GROUP ShowDatabaseResponseOwnerType
	ROLE  ShowDatabaseResponseOwnerType
}

func GetShowDatabaseResponseOwnerTypeEnum() ShowDatabaseResponseOwnerTypeEnum {
	return ShowDatabaseResponseOwnerTypeEnum{
		USER: ShowDatabaseResponseOwnerType{
			value: "USER",
		},
		GROUP: ShowDatabaseResponseOwnerType{
			value: "GROUP",
		},
		ROLE: ShowDatabaseResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c ShowDatabaseResponseOwnerType) Value() string {
	return c.value
}

func (c ShowDatabaseResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatabaseResponseOwnerType) UnmarshalJSON(b []byte) error {
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
