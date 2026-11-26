package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateConfigs **参数解释：** 域名模板配置 **约束限制：** 不涉及
type TemplateConfigs struct {

	// **参数解释：** HTTP响应头，CDN节点响应客户端时，根据HTTP响应头配置操作响应的HTTP头部信息 **约束限制：** - 调用接口修改响应头时，将覆盖原先已配置的响应头，因此，在使用此接口时，请上传全量头部信息
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	// **参数解释：** 设置源站资源在CDN节点的缓存时长 **约束限制：** 必须存在一条所有文件类型的规则
	CacheRules *[]CacheRules `json:"cache_rules,omitempty"`

	// **参数解释：** 开启回源跟随，当源站地址因业务需求做了301/302 重定向，CDN节点会先跳转到301/302对应地址获取资源，缓存后再返回给用户 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	OriginFollow302Status *string `json:"origin_follow302_status,omitempty"`

	Compress *Compress `json:"compress,omitempty"`

	// **参数解释：** Range回源，开启后，源站在收到CDN节点回源请求时，根据HTTP请求头中的Range信息返回指定范围的数据给CDN节点 **约束限制：** 开启Range回源的前提是您的源站支持Range请求，即HTTP请求头中包含Range字段，否则可能导致回源失败 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	OriginRangeStatus *string `json:"origin_range_status,omitempty"`

	IpFilter *IpFilter `json:"ip_filter,omitempty"`

	Referer *RefererConfig `json:"referer,omitempty"`

	UserAgentFilter *UserAgentFilter `json:"user_agent_filter,omitempty"`

	// **参数解释：** 设置用量封顶阈值，当实际用量大于阈值时停用域名，有效预防流量盗刷或恶意攻击带来的高额账单。  > 由于监控数据存在时延，域名将在用量达到阈值后的10分钟左右被停用  **约束限制：** 不涉及
	FlowLimitStrategy *[]FlowLimitStrategy `json:"flow_limit_strategy,omitempty"`

	// **参数解释：** 回源请求头，CDN节点回源时，在请求源站时添加指定的HTTP请求头信息 **约束限制：** - 调用接口修改回源请求头时，将覆盖原先已配置的回源请求头
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	Https *HttpTemplateBody `json:"https,omitempty"`

	Quic *Quic `json:"quic,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	HttpsTlsVersion *HttpsTlsVersion `json:"https_tls_version,omitempty"`

	// **参数解释：** 区域访问控制，通过配置访问控制规则限制客户端所在地区，实现只允许或只禁止特定地区的用户访问资源 **约束限制：** - 不涉及
	AccessAreaFilter *[]AccessAreaFilter `json:"access_area_filter,omitempty"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty"`

	// **参数解释：** 规则引擎，通过配置规则来对请求和响应进行自定义控制 **约束限制：** 不涉及
	Rules *[]RuleResponse `json:"rules,omitempty"`
}

func (o TemplateConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateConfigs struct{}"
	}

	return strings.Join([]string{"TemplateConfigs", string(data)}, " ")
}
