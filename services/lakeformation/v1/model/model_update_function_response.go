package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateFunctionResponse Response Object
type UpdateFunctionResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 函数名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	FunctionName *string `json:"function_name,omitempty"`

	// 函数类型,JAVA
	FunctionType *UpdateFunctionResponseFunctionType `json:"function_type,omitempty"`

	// 函数所有者。只能包含字母、数字和下划线，且长度为1~256个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType *UpdateFunctionResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者授权来源类型,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *UpdateFunctionResponseOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 函数类名
	ClassName *string `json:"class_name,omitempty"`

	// 创建时间格式为yyyy-mm-ddThh:mm:sss
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 函数地址信息
	ResourceUris   *[]FunctionResourceUri `json:"resource_uris,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionResponse", string(data)}, " ")
}

type UpdateFunctionResponseFunctionType struct {
	value string
}

type UpdateFunctionResponseFunctionTypeEnum struct {
	JAVA UpdateFunctionResponseFunctionType
}

func GetUpdateFunctionResponseFunctionTypeEnum() UpdateFunctionResponseFunctionTypeEnum {
	return UpdateFunctionResponseFunctionTypeEnum{
		JAVA: UpdateFunctionResponseFunctionType{
			value: "JAVA",
		},
	}
}

func (c UpdateFunctionResponseFunctionType) Value() string {
	return c.value
}

func (c UpdateFunctionResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionResponseFunctionType) UnmarshalJSON(b []byte) error {
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

type UpdateFunctionResponseOwnerType struct {
	value string
}

type UpdateFunctionResponseOwnerTypeEnum struct {
	USER  UpdateFunctionResponseOwnerType
	GROUP UpdateFunctionResponseOwnerType
	ROLE  UpdateFunctionResponseOwnerType
}

func GetUpdateFunctionResponseOwnerTypeEnum() UpdateFunctionResponseOwnerTypeEnum {
	return UpdateFunctionResponseOwnerTypeEnum{
		USER: UpdateFunctionResponseOwnerType{
			value: "USER",
		},
		GROUP: UpdateFunctionResponseOwnerType{
			value: "GROUP",
		},
		ROLE: UpdateFunctionResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c UpdateFunctionResponseOwnerType) Value() string {
	return c.value
}

func (c UpdateFunctionResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type UpdateFunctionResponseOwnerAuthSourceType struct {
	value string
}

type UpdateFunctionResponseOwnerAuthSourceTypeEnum struct {
	IAM         UpdateFunctionResponseOwnerAuthSourceType
	SAML        UpdateFunctionResponseOwnerAuthSourceType
	LDAP        UpdateFunctionResponseOwnerAuthSourceType
	LOCAL       UpdateFunctionResponseOwnerAuthSourceType
	AGENTTENANT UpdateFunctionResponseOwnerAuthSourceType
	OTHER       UpdateFunctionResponseOwnerAuthSourceType
}

func GetUpdateFunctionResponseOwnerAuthSourceTypeEnum() UpdateFunctionResponseOwnerAuthSourceTypeEnum {
	return UpdateFunctionResponseOwnerAuthSourceTypeEnum{
		IAM: UpdateFunctionResponseOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: UpdateFunctionResponseOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: UpdateFunctionResponseOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: UpdateFunctionResponseOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: UpdateFunctionResponseOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: UpdateFunctionResponseOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c UpdateFunctionResponseOwnerAuthSourceType) Value() string {
	return c.value
}

func (c UpdateFunctionResponseOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionResponseOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
