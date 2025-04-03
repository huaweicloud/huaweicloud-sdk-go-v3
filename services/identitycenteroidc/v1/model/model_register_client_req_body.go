package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RegisterClientReqBody struct {

	// 客户端名称
	ClientName string `json:"client_name"`

	// 客户端的类型。该服务仅支持public作为客户端类型
	ClientType RegisterClientReqBodyClientType `json:"client_type"`

	// 向令牌端点发送请求时所需的身份验证方法
	TokenEndpointAuthMethod RegisterClientReqBodyTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`

	// 客户端定义的作用域列表。授权后，此列表用于在授予访问令牌时限制权限
	Scopes *[]string `json:"scopes,omitempty"`

	// 客户端可以在令牌端点使用的OAuth2.0授权类型数组
	GrantTypes []RegisterClientReqBodyGrantTypes `json:"grant_types"`

	// 客户端可以在授权端点使用的OAuth2.0授权类型数组
	ResponseTypes []RegisterClientReqBodyResponseTypes `json:"response_types"`
}

func (o RegisterClientReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClientReqBody struct{}"
	}

	return strings.Join([]string{"RegisterClientReqBody", string(data)}, " ")
}

type RegisterClientReqBodyClientType struct {
	value string
}

type RegisterClientReqBodyClientTypeEnum struct {
	PUBLIC RegisterClientReqBodyClientType
}

func GetRegisterClientReqBodyClientTypeEnum() RegisterClientReqBodyClientTypeEnum {
	return RegisterClientReqBodyClientTypeEnum{
		PUBLIC: RegisterClientReqBodyClientType{
			value: "public",
		},
	}
}

func (c RegisterClientReqBodyClientType) Value() string {
	return c.value
}

func (c RegisterClientReqBodyClientType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterClientReqBodyClientType) UnmarshalJSON(b []byte) error {
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

type RegisterClientReqBodyTokenEndpointAuthMethod struct {
	value string
}

type RegisterClientReqBodyTokenEndpointAuthMethodEnum struct {
	CLIENT_SECRET_POST RegisterClientReqBodyTokenEndpointAuthMethod
}

func GetRegisterClientReqBodyTokenEndpointAuthMethodEnum() RegisterClientReqBodyTokenEndpointAuthMethodEnum {
	return RegisterClientReqBodyTokenEndpointAuthMethodEnum{
		CLIENT_SECRET_POST: RegisterClientReqBodyTokenEndpointAuthMethod{
			value: "client_secret_post",
		},
	}
}

func (c RegisterClientReqBodyTokenEndpointAuthMethod) Value() string {
	return c.value
}

func (c RegisterClientReqBodyTokenEndpointAuthMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterClientReqBodyTokenEndpointAuthMethod) UnmarshalJSON(b []byte) error {
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

type RegisterClientReqBodyGrantTypes struct {
	value string
}

type RegisterClientReqBodyGrantTypesEnum struct {
	URNIETFPARAMSOAUTHGRANT_TYPEDEVICE_CODE RegisterClientReqBodyGrantTypes
	AUTHORIZATION_CODE                      RegisterClientReqBodyGrantTypes
}

func GetRegisterClientReqBodyGrantTypesEnum() RegisterClientReqBodyGrantTypesEnum {
	return RegisterClientReqBodyGrantTypesEnum{
		URNIETFPARAMSOAUTHGRANT_TYPEDEVICE_CODE: RegisterClientReqBodyGrantTypes{
			value: "urn:ietf:params:oauth:grant-type:device_code",
		},
		AUTHORIZATION_CODE: RegisterClientReqBodyGrantTypes{
			value: "authorization_code",
		},
	}
}

func (c RegisterClientReqBodyGrantTypes) Value() string {
	return c.value
}

func (c RegisterClientReqBodyGrantTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterClientReqBodyGrantTypes) UnmarshalJSON(b []byte) error {
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

type RegisterClientReqBodyResponseTypes struct {
	value string
}

type RegisterClientReqBodyResponseTypesEnum struct {
	CODE RegisterClientReqBodyResponseTypes
}

func GetRegisterClientReqBodyResponseTypesEnum() RegisterClientReqBodyResponseTypesEnum {
	return RegisterClientReqBodyResponseTypesEnum{
		CODE: RegisterClientReqBodyResponseTypes{
			value: "code",
		},
	}
}

func (c RegisterClientReqBodyResponseTypes) Value() string {
	return c.value
}

func (c RegisterClientReqBodyResponseTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterClientReqBodyResponseTypes) UnmarshalJSON(b []byte) error {
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
