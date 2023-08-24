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
	Name string `json:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端  不支持修改
	Type AuthorizerRespType `json:"type"`

	// 只能为：FUNC
	AuthorizerType AuthorizerRespAuthorizerType `json:"authorizer_type"`

	// 函数地址。
	AuthorizerUri string `json:"authorizer_uri"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType *AuthorizerRespNetworkType `json:"network_type,omitempty"`

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

	// 自定义认证编号
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 自定义认证所属应用编号  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 自定义认证所属应用名称  暂不支持
	RomaAppName *string `json:"roma_app_name,omitempty"`
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

type AuthorizerRespAuthorizerType struct {
	value string
}

type AuthorizerRespAuthorizerTypeEnum struct {
	FUNC AuthorizerRespAuthorizerType
}

func GetAuthorizerRespAuthorizerTypeEnum() AuthorizerRespAuthorizerTypeEnum {
	return AuthorizerRespAuthorizerTypeEnum{
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

type AuthorizerRespNetworkType struct {
	value string
}

type AuthorizerRespNetworkTypeEnum struct {
	V1 AuthorizerRespNetworkType
	V2 AuthorizerRespNetworkType
}

func GetAuthorizerRespNetworkTypeEnum() AuthorizerRespNetworkTypeEnum {
	return AuthorizerRespNetworkTypeEnum{
		V1: AuthorizerRespNetworkType{
			value: "V1",
		},
		V2: AuthorizerRespNetworkType{
			value: "V2",
		},
	}
}

func (c AuthorizerRespNetworkType) Value() string {
	return c.value
}

func (c AuthorizerRespNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerRespNetworkType) UnmarshalJSON(b []byte) error {
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
