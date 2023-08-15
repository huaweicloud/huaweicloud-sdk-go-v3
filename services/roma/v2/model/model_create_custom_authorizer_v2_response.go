package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateCustomAuthorizerV2Response Response Object
type CreateCustomAuthorizerV2Response struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端
	Type CreateCustomAuthorizerV2ResponseType `json:"type"`

	// 自定义认证函数类型： - LD：自定义后端函数 - FUNC：函数服务函数
	AuthorizerType CreateCustomAuthorizerV2ResponseAuthorizerType `json:"authorizer_type"`

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

	// 自定义认证编号
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 自定义认证所属应用编号
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 自定义认证所属应用名称
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

func (c CreateCustomAuthorizerV2ResponseType) Value() string {
	return c.value
}

func (c CreateCustomAuthorizerV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCustomAuthorizerV2ResponseType) UnmarshalJSON(b []byte) error {
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

type CreateCustomAuthorizerV2ResponseAuthorizerType struct {
	value string
}

type CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum struct {
	LD   CreateCustomAuthorizerV2ResponseAuthorizerType
	FUNC CreateCustomAuthorizerV2ResponseAuthorizerType
}

func GetCreateCustomAuthorizerV2ResponseAuthorizerTypeEnum() CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum {
	return CreateCustomAuthorizerV2ResponseAuthorizerTypeEnum{
		LD: CreateCustomAuthorizerV2ResponseAuthorizerType{
			value: "LD",
		},
		FUNC: CreateCustomAuthorizerV2ResponseAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c CreateCustomAuthorizerV2ResponseAuthorizerType) Value() string {
	return c.value
}

func (c CreateCustomAuthorizerV2ResponseAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCustomAuthorizerV2ResponseAuthorizerType) UnmarshalJSON(b []byte) error {
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
