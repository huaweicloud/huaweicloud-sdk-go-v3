package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSettings struct {

	// 域名id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 网站需要登录时，设置登录页面
	LoginUrl *string `json:"login_url,omitempty" xml:"login_url"`

	// 网站需要登录时，设置登录用户名
	LoginUsername *string `json:"login_username,omitempty" xml:"login_username"`

	// 网站需要登录时，设置登录密码
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password"`

	// 网站需要登录时，设置登录cookie
	LoginCookies *string `json:"login_cookies,omitempty" xml:"login_cookies"`

	// 设置用于验证登录是否成功的网址
	VerifyUrl *string `json:"verify_url,omitempty" xml:"verify_url"`

	// 设置自定义HTTP请求头
	HttpHeaders map[string]string `json:"http_headers,omitempty" xml:"http_headers"`
}

func (o DomainSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSettings struct{}"
	}

	return strings.Join([]string{"DomainSettings", string(data)}, " ")
}
