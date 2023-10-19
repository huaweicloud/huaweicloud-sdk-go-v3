package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// FunctionInput 函数信息
type FunctionInput struct {

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	FunctionName string `json:"function_name"`

	// 函数类型，JAVA
	FunctionType FunctionInputFunctionType `json:"function_type"`

	// 函数所有者。只能包含字母、数字和下划线，且长度为1~256个字符。
	Owner string `json:"owner"`

	// 所有者类型，USER-用户,GROUP-组,ROLE-角色
	OwnerType FunctionInputOwnerType `json:"owner_type"`

	// 所有者授权来源类型,IAM-云,SAML-联邦,LDAP-ld用户,LOCAL-本地,AGENTTENANT-委托,OTHER-其它
	OwnerAuthSourceType *FunctionInputOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 函数类名。长度为1~4000个字符。
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

type FunctionInputOwnerAuthSourceType struct {
	value string
}

type FunctionInputOwnerAuthSourceTypeEnum struct {
	IAM         FunctionInputOwnerAuthSourceType
	SAML        FunctionInputOwnerAuthSourceType
	LDAP        FunctionInputOwnerAuthSourceType
	LOCAL       FunctionInputOwnerAuthSourceType
	AGENTTENANT FunctionInputOwnerAuthSourceType
	OTHER       FunctionInputOwnerAuthSourceType
}

func GetFunctionInputOwnerAuthSourceTypeEnum() FunctionInputOwnerAuthSourceTypeEnum {
	return FunctionInputOwnerAuthSourceTypeEnum{
		IAM: FunctionInputOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: FunctionInputOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: FunctionInputOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: FunctionInputOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: FunctionInputOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: FunctionInputOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c FunctionInputOwnerAuthSourceType) Value() string {
	return c.value
}

func (c FunctionInputOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionInputOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
