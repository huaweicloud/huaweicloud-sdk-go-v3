package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizerCreate struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端  不支持修改
	Type AuthorizerCreateType `json:"type"`

	// 只能为：FUNC
	AuthorizerType AuthorizerCreateAuthorizerType `json:"authorizer_type"`

	// 函数地址。
	AuthorizerUri string `json:"authorizer_uri"`

	// 函数版本。  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AuthorizerVersion *string `json:"authorizer_version,omitempty"`

	// 函数别名地址。  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AuthorizerAliasUri *string `json:"authorizer_alias_uri,omitempty"`

	// 认证来源
	Identities *[]Identity `json:"identities,omitempty"`

	// 缓存时间
	Ttl *int32 `json:"ttl,omitempty"`

	// 用户数据
	UserData *string `json:"user_data,omitempty"`

	// 自定义后端服务ID。  暂不支持
	LdApiId *string `json:"ld_api_id,omitempty"`

	// 是否发送body
	NeedBody *bool `json:"need_body,omitempty"`
}

func (o AuthorizerCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizerCreate struct{}"
	}

	return strings.Join([]string{"AuthorizerCreate", string(data)}, " ")
}

type AuthorizerCreateType struct {
	value string
}

type AuthorizerCreateTypeEnum struct {
	FRONTEND AuthorizerCreateType
	BACKEND  AuthorizerCreateType
}

func GetAuthorizerCreateTypeEnum() AuthorizerCreateTypeEnum {
	return AuthorizerCreateTypeEnum{
		FRONTEND: AuthorizerCreateType{
			value: "FRONTEND",
		},
		BACKEND: AuthorizerCreateType{
			value: "BACKEND",
		},
	}
}

func (c AuthorizerCreateType) Value() string {
	return c.value
}

func (c AuthorizerCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerCreateType) UnmarshalJSON(b []byte) error {
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

type AuthorizerCreateAuthorizerType struct {
	value string
}

type AuthorizerCreateAuthorizerTypeEnum struct {
	FUNC AuthorizerCreateAuthorizerType
}

func GetAuthorizerCreateAuthorizerTypeEnum() AuthorizerCreateAuthorizerTypeEnum {
	return AuthorizerCreateAuthorizerTypeEnum{
		FUNC: AuthorizerCreateAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c AuthorizerCreateAuthorizerType) Value() string {
	return c.value
}

func (c AuthorizerCreateAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerCreateAuthorizerType) UnmarshalJSON(b []byte) error {
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
