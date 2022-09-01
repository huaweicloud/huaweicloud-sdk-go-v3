package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSsoConfigResponse struct {

	// 是否开启SSO登录
	EnableSSO *bool `json:"enableSSO,omitempty" xml:"enableSSO"`

	// 企业域名
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// Oauth2授权地址
	AuthorizeUrl *string `json:"authorizeUrl,omitempty" xml:"authorizeUrl"`

	// 获取Token URL
	GetTokenUrl *string `json:"getTokenUrl,omitempty" xml:"getTokenUrl"`

	// APPID
	ClientId *string `json:"clientId,omitempty" xml:"clientId"`

	// APP秘钥
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret"`

	// 授权范围，OAuth2.0的OIDC取值为“openid”
	Scope *string `json:"scope,omitempty" xml:"scope"`

	// 第三方账号的字段名称
	AccFieldName *string `json:"accFieldName,omitempty" xml:"accFieldName"`

	// 用户信息查询URL
	GetUserInfoUrl *string `json:"getUserInfoUrl,omitempty" xml:"getUserInfoUrl"`

	// 鉴权类型。OAuth2.0鉴权时取0
	Oauth2ServerType *int32 `json:"oauth2ServerType,omitempty" xml:"oauth2ServerType"`

	// 拉起PC端终端的schema
	PcSchemaUrl *string `json:"pcSchemaUrl,omitempty" xml:"pcSchemaUrl"`

	// 拉起安卓端终端的schema
	AndroidSchemaUrl *string `json:"androidSchemaUrl,omitempty" xml:"androidSchemaUrl"`

	// 拉起ios端终端的schema
	IosSchemaUrl *string `json:"iosSchemaUrl,omitempty" xml:"iosSchemaUrl"`

	// 第三方姓名的字段名称
	ThirdName *string `json:"thirdName,omitempty" xml:"thirdName"`

	// 第三方邮箱的字段名称
	ThirdEmail *string `json:"thirdEmail,omitempty" xml:"thirdEmail"`

	// 第三方手机号的字段名称
	ThirdMobile *string `json:"thirdMobile,omitempty" xml:"thirdMobile"`

	// 第三方accessToken的字段名称
	ThirdAccessToken *string `json:"thirdAccessToken,omitempty" xml:"thirdAccessToken"`

	// 第三方头像链接的字段名称
	ThirdHeadImgUrl *string `json:"thirdHeadImgUrl,omitempty" xml:"thirdHeadImgUrl"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowSsoConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSsoConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSsoConfigResponse", string(data)}, " ")
}
