package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Actions **参数解释：** 规则行为 **约束限制：** 不涉及
type Actions struct {

	// **参数解释：** 高级回源，实现根据不同的资源类型或路径回源到不同源站 **约束限制：** 最多配置20条
	FlexibleOrigin *[]FlexibleOriginsEngine `json:"flexible_origin,omitempty"`

	// **参数解释：** CDN节点回源时，改写用户回源请求URL的HTTP头部信息 **约束限制：** - 该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息 - 如果域名在后台配置了特殊请求头，需要将对应的请求头一并传入
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	// **参数解释：** 配置节点响应给客户端的头部信息，配置响应消息后，用户请求加速域名下的资源时，CDN返回给用户的消息中将包含该域名配置的响应头信息 **约束限制：** - 该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息 - 如果域名在后台配置了特殊请求头，需要将对应的请求头一并传入
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	AccessControl *AccessControl `json:"access_control,omitempty"`

	RequestLimitRule *RequestLimitRulesEngine `json:"request_limit_rule,omitempty"`

	OriginRequestUrlRewrite *OriginRequestUrlRewriteEngine `json:"origin_request_url_rewrite,omitempty"`

	CacheRule *CacheRulesEngine `json:"cache_rule,omitempty"`

	RequestUrlRewrite *RequestUrlRewriteEngine `json:"request_url_rewrite,omitempty"`

	BrowserCacheRule *BrowserCacheRulesEngine `json:"browser_cache_rule,omitempty"`

	ErrorCodeCache *[]ErrorCodeCacheEngine `json:"error_code_cache,omitempty"`
}

func (o Actions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Actions struct{}"
	}

	return strings.Join([]string{"Actions", string(data)}, " ")
}
