package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateFunctionResponse Response Object
type CreateFunctionResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	FunctionName *string `json:"function_name,omitempty"`

	// 函数类型,JAVA
	FunctionType *CreateFunctionResponseFunctionType `json:"function_type,omitempty"`

	// 函数所有者。只能包含字母、数字和下划线，且长度为1~256个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType *CreateFunctionResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者授权来源类型,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *CreateFunctionResponseOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 函数类名
	ClassName *string `json:"class_name,omitempty"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris   *[]FunctionResourceUri `json:"resource_uris,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionResponse struct{}"
	}

	return strings.Join([]string{"CreateFunctionResponse", string(data)}, " ")
}

type CreateFunctionResponseFunctionType struct {
	value string
}

type CreateFunctionResponseFunctionTypeEnum struct {
	JAVA CreateFunctionResponseFunctionType
}

func GetCreateFunctionResponseFunctionTypeEnum() CreateFunctionResponseFunctionTypeEnum {
	return CreateFunctionResponseFunctionTypeEnum{
		JAVA: CreateFunctionResponseFunctionType{
			value: "JAVA",
		},
	}
}

func (c CreateFunctionResponseFunctionType) Value() string {
	return c.value
}

func (c CreateFunctionResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionResponseFunctionType) UnmarshalJSON(b []byte) error {
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

type CreateFunctionResponseOwnerType struct {
	value string
}

type CreateFunctionResponseOwnerTypeEnum struct {
	USER  CreateFunctionResponseOwnerType
	GROUP CreateFunctionResponseOwnerType
	ROLE  CreateFunctionResponseOwnerType
}

func GetCreateFunctionResponseOwnerTypeEnum() CreateFunctionResponseOwnerTypeEnum {
	return CreateFunctionResponseOwnerTypeEnum{
		USER: CreateFunctionResponseOwnerType{
			value: "USER",
		},
		GROUP: CreateFunctionResponseOwnerType{
			value: "GROUP",
		},
		ROLE: CreateFunctionResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c CreateFunctionResponseOwnerType) Value() string {
	return c.value
}

func (c CreateFunctionResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type CreateFunctionResponseOwnerAuthSourceType struct {
	value string
}

type CreateFunctionResponseOwnerAuthSourceTypeEnum struct {
	IAM         CreateFunctionResponseOwnerAuthSourceType
	SAML        CreateFunctionResponseOwnerAuthSourceType
	LDAP        CreateFunctionResponseOwnerAuthSourceType
	LOCAL       CreateFunctionResponseOwnerAuthSourceType
	AGENTTENANT CreateFunctionResponseOwnerAuthSourceType
	OTHER       CreateFunctionResponseOwnerAuthSourceType
}

func GetCreateFunctionResponseOwnerAuthSourceTypeEnum() CreateFunctionResponseOwnerAuthSourceTypeEnum {
	return CreateFunctionResponseOwnerAuthSourceTypeEnum{
		IAM: CreateFunctionResponseOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: CreateFunctionResponseOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: CreateFunctionResponseOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: CreateFunctionResponseOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: CreateFunctionResponseOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: CreateFunctionResponseOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c CreateFunctionResponseOwnerAuthSourceType) Value() string {
	return c.value
}

func (c CreateFunctionResponseOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionResponseOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
