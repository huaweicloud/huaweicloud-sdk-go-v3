package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthorizeConfigInfoRequestBody struct {
	// 是否开启SSO登录

	EnableSSO bool `json:"enableSSO"`
	// 企业域名 说明：开启SSO登录时必填

	Domain *string `json:"domain,omitempty"`
	// Oauth2授权地址 说明：开启SSO登录时必填

	AuthorizeUrl *string `json:"authorizeUrl,omitempty"`
	// 获取Token URL 说明：开启SSO登录时必填

	GetTokenUrl *string `json:"getTokenUrl,omitempty"`
	// APPID 说明：开启SSO登录时必填

	ClientId *string `json:"clientId,omitempty"`
	// APP秘钥 说明：开启SSO登录时，若不修改APP秘钥，则置空即可

	ClientSecret *string `json:"clientSecret,omitempty"`
	// 授权范围，OAuth2.0的OIDC取值为“openid”

	Scope *string `json:"scope,omitempty"`
	// 第三方账号的字段名称 说明：开启SSO登录时必填

	AccFieldName *string `json:"accFieldName,omitempty"`
	// 用户信息查询URL

	GetUserInfoUrl *string `json:"getUserInfoUrl,omitempty"`
	// 鉴权类型。OAuth2.0鉴权时取0

	Oauth2ServerType *int32 `json:"oauth2ServerType,omitempty"`
	// 拉起PC端终端的schema

	PcSchemaUrl *string `json:"pcSchemaUrl,omitempty"`
	// 拉起安卓端终端的schema

	AndroidSchemaUrl *string `json:"androidSchemaUrl,omitempty"`
	// 拉起ios端终端的schema

	IosSchemaUrl *string `json:"iosSchemaUrl,omitempty"`
	// 第三方姓名的字段名称

	ThirdName *string `json:"thirdName,omitempty"`
	// 第三方邮箱的字段名称

	ThirdEmail *string `json:"thirdEmail,omitempty"`
	// 第三方手机号的字段名称

	ThirdMobile *string `json:"thirdMobile,omitempty"`
	// 第三方accessToken的字段名称 说明：开启SSO登录时必填

	ThirdAccessToken *string `json:"thirdAccessToken,omitempty"`
	// 第三方头像链接的字段名称

	ThirdHeadImgUrl *string `json:"thirdHeadImgUrl,omitempty"`
}

func (o AuthorizeConfigInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeConfigInfoRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeConfigInfoRequestBody", string(data)}, " ")
}
