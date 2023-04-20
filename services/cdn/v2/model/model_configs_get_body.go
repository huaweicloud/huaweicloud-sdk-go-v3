package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type ConfigsGetBody struct {

	// 回源请求头配置
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	// http header配置
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuthGetBody `json:"url_auth,omitempty"`

	Https *HttpGetBody `json:"https,omitempty"`

	Sources *[]SourcesConfig `json:"sources,omitempty"`

	OriginProtocol *string `json:"origin_protocol,omitempty"`

	// 回源跟随（on：开启，off：关闭）。
	OriginFollow302Status *string `json:"origin_follow302_status,omitempty"`

	CacheRules *[]CacheRules `json:"cache_rules,omitempty"`

	IpFilter *IpFilter `json:"ip_filter,omitempty"`

	Referer *RefererConfig `json:"referer,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	Compress *Compress `json:"compress,omitempty"`

	CacheUrlParameterFilter *CacheUrlParameterFilter `json:"cache_url_parameter_filter,omitempty"`

	ErrorCodeCache *[]ErrorCodeCache `json:"error_code_cache,omitempty"`

	UserAgentFilter *UserAgentFilter `json:"user_agent_filter,omitempty"`

	// 改写回源URL，最多配置20条。
	OriginRequestUrlRewrite *[]OriginRequestUrlRewrite `json:"origin_request_url_rewrite,omitempty"`

	// 自定义错误页面
	ErrorCodeRedirectRules *[]ErrorCodeRedirectRules `json:"error_code_redirect_rules,omitempty"`
}

func (o ConfigsGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsGetBody struct{}"
	}

	return strings.Join([]string{"ConfigsGetBody", string(data)}, " ")
}
