package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceInput ResourceInput
type ResourceInput struct {

	// 元数据资源类型,CATALOG,DATABASE,TABLE,FUNC,MODEL,COLUMN,URI
	ResourceType ResourceInputResourceType `json:"resource_type"`

	// catalog名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	Catalog *string `json:"catalog,omitempty"`

	// 数据库名称。只能包含中文、字母、数字和下划线，且长度为1到128个字符。
	Database *string `json:"database,omitempty"`

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	Function *string `json:"function,omitempty"`

	// 表名称。只能包含中文、字母、数字和下划线，且长度为1~256个字符。
	Table *string `json:"table,omitempty"`

	// 列名称。只能包含中文、字母、数字和_-+*\\(), 特殊字符，且长度为1~767个字符。
	Column *string `json:"column,omitempty"`

	// URI
	Uri *string `json:"uri,omitempty"`

	// 列名称列表
	Columns *[]string `json:"columns,omitempty"`
}

func (o ResourceInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInput struct{}"
	}

	return strings.Join([]string{"ResourceInput", string(data)}, " ")
}

type ResourceInputResourceType struct {
	value string
}

type ResourceInputResourceTypeEnum struct {
	CATALOG  ResourceInputResourceType
	DATABASE ResourceInputResourceType
	TABLE    ResourceInputResourceType
	FUNC     ResourceInputResourceType
	MODEL    ResourceInputResourceType
	COLUMN   ResourceInputResourceType
	URI      ResourceInputResourceType
}

func GetResourceInputResourceTypeEnum() ResourceInputResourceTypeEnum {
	return ResourceInputResourceTypeEnum{
		CATALOG: ResourceInputResourceType{
			value: "CATALOG",
		},
		DATABASE: ResourceInputResourceType{
			value: "DATABASE",
		},
		TABLE: ResourceInputResourceType{
			value: "TABLE",
		},
		FUNC: ResourceInputResourceType{
			value: "FUNC",
		},
		MODEL: ResourceInputResourceType{
			value: "MODEL",
		},
		COLUMN: ResourceInputResourceType{
			value: "COLUMN",
		},
		URI: ResourceInputResourceType{
			value: "URI",
		},
	}
}

func (c ResourceInputResourceType) Value() string {
	return c.value
}

func (c ResourceInputResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceInputResourceType) UnmarshalJSON(b []byte) error {
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
