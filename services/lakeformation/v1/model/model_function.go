package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 函数信息
type Function struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 函数名
	FunctionName string `json:"function_name"`

	// 函数类型
	FunctionType FunctionFunctionType `json:"function_type"`

	// 函数所有者
	Owner string `json:"owner"`

	// 所有者类型
	OwnerType FunctionOwnerType `json:"owner_type"`

	// 函数类名
	ClassName string `json:"class_name"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris *[]FunctionResourceUri `json:"resource_uris,omitempty"`
}

func (o Function) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Function struct{}"
	}

	return strings.Join([]string{"Function", string(data)}, " ")
}

type FunctionFunctionType struct {
	value string
}

type FunctionFunctionTypeEnum struct {
	JAVA FunctionFunctionType
}

func GetFunctionFunctionTypeEnum() FunctionFunctionTypeEnum {
	return FunctionFunctionTypeEnum{
		JAVA: FunctionFunctionType{
			value: "JAVA",
		},
	}
}

func (c FunctionFunctionType) Value() string {
	return c.value
}

func (c FunctionFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionFunctionType) UnmarshalJSON(b []byte) error {
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

type FunctionOwnerType struct {
	value string
}

type FunctionOwnerTypeEnum struct {
	USER  FunctionOwnerType
	GROUP FunctionOwnerType
	ROLE  FunctionOwnerType
}

func GetFunctionOwnerTypeEnum() FunctionOwnerTypeEnum {
	return FunctionOwnerTypeEnum{
		USER: FunctionOwnerType{
			value: "USER",
		},
		GROUP: FunctionOwnerType{
			value: "GROUP",
		},
		ROLE: FunctionOwnerType{
			value: "ROLE",
		},
	}
}

func (c FunctionOwnerType) Value() string {
	return c.value
}

func (c FunctionOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionOwnerType) UnmarshalJSON(b []byte) error {
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
