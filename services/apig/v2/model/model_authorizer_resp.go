package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
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
	FUNC AuthorizerRespAuthorizerType
}

func GetAuthorizerRespAuthorizerTypeEnum() AuthorizerRespAuthorizerTypeEnum {
	return AuthorizerRespAuthorizerTypeEnum{
		FUNC: AuthorizerRespAuthorizerType{
			value: "FUNC",
		},
	}
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
