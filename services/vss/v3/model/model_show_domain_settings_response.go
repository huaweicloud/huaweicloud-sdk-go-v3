package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainSettingsResponse Response Object
type ShowDomainSettingsResponse struct {

	// 网站域名ID
	DomainId *string `json:"domain_id,omitempty"`

	// 网站需要登录时，设置登录页面
	LoginUrl *string `json:"login_url,omitempty"`

	// 网站需要登录时，设置登录用户名
	LoginUsername *string `json:"login_username,omitempty"`

	// 网站需要登录时，设置登录密码
	LoginPassword *string `json:"login_password,omitempty"`

	// 网站需要登录时，设置登录cookie
	LoginCookies *string `json:"login_cookies,omitempty"`

	// 设置用于验证登录是否成功的网址
	VerifyUrl *string `json:"verify_url,omitempty"`

	// 设置自定义HTTP请求头
	HttpHeaders map[string]string `json:"http_headers,omitempty"`

	// 网站域名
	DomainName     *string `json:"domain_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainSettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainSettingsResponse", string(data)}, " ")
}
