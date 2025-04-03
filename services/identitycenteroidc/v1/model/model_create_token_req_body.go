package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTokenReqBody struct {

	// 客户端的唯一标识
	ClientId string `json:"client_id"`

	// 为客户端生成的秘密字符串。客户端将使用此字符串在后续调用中获得服务的身份验证
	ClientSecret string `json:"client_secret"`

	// 从授权服务接收的授权代码。执行授权授予请求以获取对令牌的访问权限时需要此参数
	Code *string `json:"code,omitempty"`

	// 仅在为设备代码授权类型调用此API时使用
	DeviceCode *string `json:"device_code,omitempty"`

	// 请求的授权类型。支持授权码、设备代码、客户端凭证和刷新令牌等授权类型
	GrantType CreateTokenReqBodyGrantType `json:"grant_type"`

	// 将接收授权代码的应用程序的位置。用户授权服务将请求发送到此位置
	RedirectUri *string `json:"redirect_uri,omitempty"`

	// 刷新令牌，此令牌可在访问令牌过期后获取新的访问令牌
	RefreshToken *string `json:"refresh_token,omitempty"`

	// 客户端定义的作用域列表，表示客户端想要获取的权限。授权后，此列表用于在授予访问令牌时限制权限
	Scopes *[]string `json:"scopes,omitempty"`
}

func (o CreateTokenReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTokenReqBody struct{}"
	}

	return strings.Join([]string{"CreateTokenReqBody", string(data)}, " ")
}

type CreateTokenReqBodyGrantType struct {
	value string
}

type CreateTokenReqBodyGrantTypeEnum struct {
	AUTHORIZATION_CODE                      CreateTokenReqBodyGrantType
	URNIETFPARAMSOAUTHGRANT_TYPEDEVICE_CODE CreateTokenReqBodyGrantType
}

func GetCreateTokenReqBodyGrantTypeEnum() CreateTokenReqBodyGrantTypeEnum {
	return CreateTokenReqBodyGrantTypeEnum{
		AUTHORIZATION_CODE: CreateTokenReqBodyGrantType{
			value: "authorization_code",
		},
		URNIETFPARAMSOAUTHGRANT_TYPEDEVICE_CODE: CreateTokenReqBodyGrantType{
			value: "urn:ietf:params:oauth:grant-type:device_code",
		},
	}
}

func (c CreateTokenReqBodyGrantType) Value() string {
	return c.value
}

func (c CreateTokenReqBodyGrantType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTokenReqBodyGrantType) UnmarshalJSON(b []byte) error {
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
