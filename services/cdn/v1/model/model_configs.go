package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type Configs struct {

	// 回源请求头改写 该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息。
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty" xml:"origin_request_header"`

	// http header配置 该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息。
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty" xml:"http_response_header"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty" xml:"url_auth"`

	Https *HttpPutBody `json:"https,omitempty" xml:"https"`

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

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
