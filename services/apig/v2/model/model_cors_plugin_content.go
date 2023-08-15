package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 跨域资源共享插件类型
type CorsPluginContent struct {

	// Access-Control-Allow-Origin头，该字段必填，允许访问该资源的外域URI。对于不需要携带身份凭证的请求，服务器可以指定该字段的值为通配符*，表示允许来自所有域的请求。 多个域名使用英文逗号分隔。
	AllowOrigin string `json:"allow_origin"`

	// Access-Control-Allow-Methods头，请求所允许使用的 HTTP 方法。 多个方法使用英文逗号分隔。
	AllowMethods *string `json:"allow_methods,omitempty"`

	// Access-Control-Allow-Headers头，请求中允许携带的头域字段。 多个头域使用英文逗号分隔。
	AllowHeaders *string `json:"allow_headers,omitempty"`

	// Access-Control-Expose-Headers 头，让服务器把允许浏览器访问的头放入白名单。 多个头域可通过英文逗号分隔。
	ExposeHeaders *string `json:"expose_headers,omitempty"`

	// Access-Control-Max-Age 头，表示本次预检的有效期，单位：秒,范围为0-86400。在有效期内，无需再次发出预检请求。
	MaxAge *int32 `json:"max_age,omitempty"`

	// Access-Control-Allow-Credentials 头，是否允许浏览器读取response的内容。
	AllowCredentials *bool `json:"allow_credentials,omitempty"`
}

func (o CorsPluginContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorsPluginContent struct{}"
	}

	return strings.Join([]string{"CorsPluginContent", string(data)}, " ")
}
