package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type ConfigsGetBody struct {

	// 回源请求头配置
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty" xml:"origin_request_header"`

	// http header配置
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty" xml:"http_response_header"`

	UrlAuth *UrlAuthGetBody `json:"url_auth,omitempty" xml:"url_auth"`

	Https *HttpGetBody `json:"https,omitempty" xml:"https"`

	// 源站配置。
	Sources *[]SourcesConfig `json:"sources,omitempty" xml:"sources"`

	// 回源协议（follow：协议跟随回源，http：HTTP回源(默认)，https：https回源）。
	OriginProtocol *string `json:"origin_protocol,omitempty" xml:"origin_protocol"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty" xml:"force_redirect"`

	Compress *Compress `json:"compress,omitempty" xml:"compress"`

	CacheUrlParameterFilter *CacheUrlParameterFilter `json:"cache_url_parameter_filter,omitempty" xml:"cache_url_parameter_filter"`

	// ipv6设置（1：打开；0：关闭）
	Ipv6Accelerate *int32 `json:"ipv6_accelerate,omitempty" xml:"ipv6_accelerate"`
}

func (o ConfigsGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsGetBody struct{}"
	}

	return strings.Join([]string{"ConfigsGetBody", string(data)}, " ")
}
