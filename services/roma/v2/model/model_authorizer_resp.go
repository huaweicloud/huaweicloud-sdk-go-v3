package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type AuthorizerResp struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name" xml:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端
	Type AuthorizerRespType `json:"type" xml:"type"`

	// 自定义认证函数类型： - LD：自定义后端函数 - FUNC：函数服务函数
	AuthorizerType AuthorizerRespAuthorizerType `json:"authorizer_type" xml:"authorizer_type"`

	// 函数地址。  注意：使用自定义后端的函数API，API请求方法必须为POST，且API状态必须为已部署。
	AuthorizerUri string `json:"authorizer_uri" xml:"authorizer_uri"`

	// 认证来源
	Identities *[]Identity `json:"identities,omitempty" xml:"identities"`

	// 缓存时间
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 用户数据
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 自定义后端服务ID。  自定义认证函数类型为LD时必填
	LdApiId *string `json:"ld_api_id,omitempty" xml:"ld_api_id"`

	// 是否发送body
	NeedBody *bool `json:"need_body,omitempty" xml:"need_body"`

	// 自定义认证编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 自定义认证所属应用编号
	RomaAppId *string `json:"roma_app_id,omitempty" xml:"roma_app_id"`

	// 自定义认证所属应用名称
	RomaAppName *string `json:"roma_app_name,omitempty" xml:"roma_app_name"`
}

func (o AuthorizerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizerResp struct{}"
	}

	return strings.Join([]string{"AuthorizerResp", string(data)}, " ")
}

type AuthorizerRespType struct {
	value string
}

type AuthorizerRespTypeEnum struct {
	FRONTEND AuthorizerRespType
	BACKEND  AuthorizerRespType
}

func GetAuthorizerRespTypeEnum() AuthorizerRespTypeEnum {
	return AuthorizerRespTypeEnum{
		FRONTEND: AuthorizerRespType{
			value: "FRONTEND",
		},
		BACKEND: AuthorizerRespType{
			value: "BACKEND",
		},
	}
}

func (c AuthorizerRespType) Value() string {
	return c.value
}

func (c AuthorizerRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerRespType) UnmarshalJSON(b []byte) error {
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

type AuthorizerRespAuthorizerType struct {
	value string
}

type AuthorizerRespAuthorizerTypeEnum struct {
	LD   AuthorizerRespAuthorizerType
	FUNC AuthorizerRespAuthorizerType
}

func GetAuthorizerRespAuthorizerTypeEnum() AuthorizerRespAuthorizerTypeEnum {
	return AuthorizerRespAuthorizerTypeEnum{
		LD: AuthorizerRespAuthorizerType{
			value: "LD",
		},
		FUNC: AuthorizerRespAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c AuthorizerRespAuthorizerType) Value() string {
	return c.value
}

func (c AuthorizerRespAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerRespAuthorizerType) UnmarshalJSON(b []byte) error {
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
