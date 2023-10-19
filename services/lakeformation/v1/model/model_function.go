package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Function 函数信息
type Function struct {

	// catalog名称
	CatalogName string `json:"catalog_name"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	FunctionName string `json:"function_name"`

	// 函数类型,JAVA
	FunctionType FunctionFunctionType `json:"function_type"`

	// 函数所有者。只能包含字母、数字和下划线，且长度为1~256个字符。
	Owner string `json:"owner"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType FunctionOwnerType `json:"owner_type"`

	// 所有者授权来源类型,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *FunctionOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

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

type FunctionOwnerAuthSourceType struct {
	value string
}

type FunctionOwnerAuthSourceTypeEnum struct {
	IAM         FunctionOwnerAuthSourceType
	SAML        FunctionOwnerAuthSourceType
	LDAP        FunctionOwnerAuthSourceType
	LOCAL       FunctionOwnerAuthSourceType
	AGENTTENANT FunctionOwnerAuthSourceType
	OTHER       FunctionOwnerAuthSourceType
}

func GetFunctionOwnerAuthSourceTypeEnum() FunctionOwnerAuthSourceTypeEnum {
	return FunctionOwnerAuthSourceTypeEnum{
		IAM: FunctionOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: FunctionOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: FunctionOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: FunctionOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: FunctionOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: FunctionOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c FunctionOwnerAuthSourceType) Value() string {
	return c.value
}

func (c FunctionOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
