package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type AuthorizerBase struct {
	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。

	Name string `json:"name"`
	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端

	Type AuthorizerBaseType `json:"type"`
	// 自定义认证函数类型： - LD：自定义后端函数 - FUNC：函数服务函数

	AuthorizerType AuthorizerBaseAuthorizerType `json:"authorizer_type"`
	// 函数地址。  注意：使用自定义后端的函数API，API请求方法必须为POST，且API状态必须为已部署。

	AuthorizerUri string `json:"authorizer_uri"`
	// 认证来源

	Identities *[]Identity `json:"identities,omitempty"`
	// 缓存时间

	Ttl *int32 `json:"ttl,omitempty"`
	// 用户数据

	UserData *string `json:"user_data,omitempty"`
	// 自定义后端服务ID。  自定义认证函数类型为LD时必填

	LdApiId *string `json:"ld_api_id,omitempty"`
	// 是否发送body

	NeedBody *bool `json:"need_body,omitempty"`
}

func (o AuthorizerBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizerBase struct{}"
	}

	return strings.Join([]string{"AuthorizerBase", string(data)}, " ")
}

type AuthorizerBaseType struct {
	value string
}

type AuthorizerBaseTypeEnum struct {
	FRONTEND AuthorizerBaseType
	BACKEND  AuthorizerBaseType
}

func GetAuthorizerBaseTypeEnum() AuthorizerBaseTypeEnum {
	return AuthorizerBaseTypeEnum{
		FRONTEND: AuthorizerBaseType{
			value: "FRONTEND",
		},
		BACKEND: AuthorizerBaseType{
			value: "BACKEND",
		},
	}
}

func (c AuthorizerBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerBaseType) UnmarshalJSON(b []byte) error {
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

type AuthorizerBaseAuthorizerType struct {
	value string
}

type AuthorizerBaseAuthorizerTypeEnum struct {
	LD   AuthorizerBaseAuthorizerType
	FUNC AuthorizerBaseAuthorizerType
}

func GetAuthorizerBaseAuthorizerTypeEnum() AuthorizerBaseAuthorizerTypeEnum {
	return AuthorizerBaseAuthorizerTypeEnum{
		LD: AuthorizerBaseAuthorizerType{
			value: "LD",
		},
		FUNC: AuthorizerBaseAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c AuthorizerBaseAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerBaseAuthorizerType) UnmarshalJSON(b []byte) error {
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
