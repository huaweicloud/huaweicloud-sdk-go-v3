package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CorsConfig 参数解释：转发策略跨域的配置。
type CorsConfig struct {

	// 参数解释：允许的访问来源列表。支持只配置一个元素*，或配置一个或多个值。  约束限制： - 单个值必须以http://或者https://开头，后边加一个正确的域名或一级泛域名。（例：http://_*.test.abc.example.com） - 单个值可以不加端口，也可以指定端口，端口范围：1~65535。
	AllowOrigin *[]string `json:"allow_origin,omitempty"`

	// 参数解释：选择跨域访问时允许的 HTTP 方法。
	AllowMethods *[]string `json:"allow_methods,omitempty"`

	// 参数解释：允许跨域的 Header 列表。
	AllowHeaders *[]string `json:"allow_headers,omitempty"`

	// 参数解释：允许暴露的Header列表。
	ExposeHeaders *[]string `json:"expose_headers,omitempty"`

	// 参数解释：是否允许携带凭证信息。  取值范围： - true：是。 - false：否。
	AllowCredentials *bool `json:"allow_credentials,omitempty"`

	// 参数解释：预检请求在浏览器的最大缓存时间，单位：秒。  取值范围：-1~172800。
	MaxAge *int64 `json:"max_age,omitempty"`
}

func (o CorsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorsConfig struct{}"
	}

	return strings.Join([]string{"CorsConfig", string(data)}, " ")
}
