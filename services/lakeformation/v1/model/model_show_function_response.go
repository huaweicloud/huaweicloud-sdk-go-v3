package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowFunctionResponse Response Object
type ShowFunctionResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	FunctionName *string `json:"function_name,omitempty"`

	// 函数类型,JAVA
	FunctionType *ShowFunctionResponseFunctionType `json:"function_type,omitempty"`

	// 函数所有者。只能包含字母、数字和下划线，且长度为1~256个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType *ShowFunctionResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者授权来源类型,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *ShowFunctionResponseOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 函数类名
	ClassName *string `json:"class_name,omitempty"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris   *[]FunctionResourceUri `json:"resource_uris,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionResponse", string(data)}, " ")
}

type ShowFunctionResponseFunctionType struct {
	value string
}

type ShowFunctionResponseFunctionTypeEnum struct {
	JAVA ShowFunctionResponseFunctionType
}

func GetShowFunctionResponseFunctionTypeEnum() ShowFunctionResponseFunctionTypeEnum {
	return ShowFunctionResponseFunctionTypeEnum{
		JAVA: ShowFunctionResponseFunctionType{
			value: "JAVA",
		},
	}
}

func (c ShowFunctionResponseFunctionType) Value() string {
	return c.value
}

func (c ShowFunctionResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionResponseFunctionType) UnmarshalJSON(b []byte) error {
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

type ShowFunctionResponseOwnerType struct {
	value string
}

type ShowFunctionResponseOwnerTypeEnum struct {
	USER  ShowFunctionResponseOwnerType
	GROUP ShowFunctionResponseOwnerType
	ROLE  ShowFunctionResponseOwnerType
}

func GetShowFunctionResponseOwnerTypeEnum() ShowFunctionResponseOwnerTypeEnum {
	return ShowFunctionResponseOwnerTypeEnum{
		USER: ShowFunctionResponseOwnerType{
			value: "USER",
		},
		GROUP: ShowFunctionResponseOwnerType{
			value: "GROUP",
		},
		ROLE: ShowFunctionResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c ShowFunctionResponseOwnerType) Value() string {
	return c.value
}

func (c ShowFunctionResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type ShowFunctionResponseOwnerAuthSourceType struct {
	value string
}

type ShowFunctionResponseOwnerAuthSourceTypeEnum struct {
	IAM         ShowFunctionResponseOwnerAuthSourceType
	SAML        ShowFunctionResponseOwnerAuthSourceType
	LDAP        ShowFunctionResponseOwnerAuthSourceType
	LOCAL       ShowFunctionResponseOwnerAuthSourceType
	AGENTTENANT ShowFunctionResponseOwnerAuthSourceType
	OTHER       ShowFunctionResponseOwnerAuthSourceType
}

func GetShowFunctionResponseOwnerAuthSourceTypeEnum() ShowFunctionResponseOwnerAuthSourceTypeEnum {
	return ShowFunctionResponseOwnerAuthSourceTypeEnum{
		IAM: ShowFunctionResponseOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: ShowFunctionResponseOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: ShowFunctionResponseOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: ShowFunctionResponseOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: ShowFunctionResponseOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: ShowFunctionResponseOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c ShowFunctionResponseOwnerAuthSourceType) Value() string {
	return c.value
}

func (c ShowFunctionResponseOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionResponseOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
