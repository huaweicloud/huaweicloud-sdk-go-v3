package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 函数信息
type FunctionInput struct {

	// 函数名
	FunctionName string `json:"function_name"`

	// 函数类型
	FunctionType FunctionInputFunctionType `json:"function_type"`

	// 函数所有者
	Owner string `json:"owner"`

	// 所有者类型
	OwnerType FunctionInputOwnerType `json:"owner_type"`

	// 函数类名
	ClassName string `json:"class_name"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris *[]FunctionResourceUri `json:"resource_uris,omitempty"`
}

func (o FunctionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionInput struct{}"
	}

	return strings.Join([]string{"FunctionInput", string(data)}, " ")
}

type FunctionInputFunctionType struct {
	value string
}

type FunctionInputFunctionTypeEnum struct {
	JAVA FunctionInputFunctionType
}

func GetFunctionInputFunctionTypeEnum() FunctionInputFunctionTypeEnum {
	return FunctionInputFunctionTypeEnum{
		JAVA: FunctionInputFunctionType{
			value: "JAVA",
		},
	}
}

func (c FunctionInputFunctionType) Value() string {
	return c.value
}

func (c FunctionInputFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionInputFunctionType) UnmarshalJSON(b []byte) error {
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

type FunctionInputOwnerType struct {
	value string
}

type FunctionInputOwnerTypeEnum struct {
	USER  FunctionInputOwnerType
	GROUP FunctionInputOwnerType
	ROLE  FunctionInputOwnerType
}

func GetFunctionInputOwnerTypeEnum() FunctionInputOwnerTypeEnum {
	return FunctionInputOwnerTypeEnum{
		USER: FunctionInputOwnerType{
			value: "USER",
		},
		GROUP: FunctionInputOwnerType{
			value: "GROUP",
		},
		ROLE: FunctionInputOwnerType{
			value: "ROLE",
		},
	}
}

func (c FunctionInputOwnerType) Value() string {
	return c.value
}

func (c FunctionInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionInputOwnerType) UnmarshalJSON(b []byte) error {
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
