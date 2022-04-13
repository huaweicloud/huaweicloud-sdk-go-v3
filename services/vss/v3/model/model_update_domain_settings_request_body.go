package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainSettingsRequestBody struct {
	// 域名id

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
}

func (o UpdateDomainSettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSettingsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainSettingsRequestBody", string(data)}, " ")
}
