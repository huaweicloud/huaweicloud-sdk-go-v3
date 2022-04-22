package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateCustomAuthorizerV2Response struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端  不支持修改
	Type CreateCustomAuthorizerV2ResponseType `json:"type"`

	// 只能为：FUNC
	AuthorizerType CreateCustomAuthorizerV2ResponseAuthorizerType `json:"authorizer_type"`

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
	RomaAppName    *string `json:"roma_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomAuthorizerV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomAuthorizerV2Response struct{}"
	}

	return strings.Join([]string{"CreateCustomAuthorizerV2Response", string(data)}, " ")
}

type CreateCustomAuthorizerV2ResponseType struct {
	value string
}

type CreateCustomAuthorizerV2ResponseTypeEnum struct {
	FRONTEND CreateCustomAuthorizerV2ResponseType
	BACKEND  CreateCustomAuthorizerV2ResponseType
}

func GetCreateCustomAuthorizerV2ResponseTypeEnum() CreateCustomAuthorizerV2ResponseTypeEnum {
	return CreateCustomAuthorizerV2ResponseTypeEnum{
		FRONTEND: CreateCustomAuthorizerV2ResponseType{
			value: "FRONTEND",
		},
		BACKEND: CreateCustomAuthorizerV2ResponseType{
			value: "BACKEND",
		},
	}
}

func (c CreateCustomAuthorizerV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCustomAuthorizerV2ResponseType) UnmarshalJSON(b []byte) error {
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

type CreateCustomAuthorizerV2ResponseAuthorizerType struct {
	value string
}

type CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum struct {
	FUNC CreateCustomAuthorizerV2ResponseAuthorizerType
}

func GetCreateCustomAuthorizerV2ResponseAuthorizerTypeEnum() CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum {
	return CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum{
		FUNC: CreateCustomAuthorizerV2ResponseAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c CreateCustomAuthorizerV2ResponseAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCustomAuthorizerV2ResponseAuthorizerType) UnmarshalJSON(b []byte) error {
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
