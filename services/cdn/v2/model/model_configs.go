package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Configs **参数解释：** 配置项 **约束限制：** 不涉及
type Configs struct {

	// **参数解释：** 业务类型 **约束限制：** 暂不支持“全站加速”变更为其它业务类型 **取值范围：** - web: 网站加速 - download: 文件下载加速 - video: 点播加速  **默认取值：** 不涉及
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释：** 加速域名的服务范围，客户端访问将会调度到服务范围内的CDN节点，可根据用户聚集区域选择合适的范围 **约束限制：** - 暂不支持“中国大陆”与“中国大陆境外”互相直接切换 - 服务范围修改为中国大陆或全球时，加速域名需要到工信部备案  **取值范围：** - mainland_china：中国大陆 - global：全球 - outside_mainland_china：中国大陆境外  **默认取值：** 不涉及
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 域名备注信息 **约束限制：** 不涉及 **取值范围：** 0-10000个字符 **默认取值：** 不涉及
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 回源请求头，CDN节点回源时，根据回源请求头配置操作请求的HTTP头部信息 **约束限制：** - 修改时该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息 - 如果域名在后台配置了特殊请求头，需要将对应的请求头一并传入
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	// **参数解释：** HTTP响应头，CDN节点响应客户端时，根据HTTP响应头配置操作响应的HTTP头部信息 **约束限制：** - 修改时该功能将覆盖原有配置（清空之前的配置），在使用此接口时，请上传全量头部信息 - 如果域名在后台配置了特殊响应头，需要将对应的响应头一并传入
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty"`

	Https *HttpPutBody `json:"https,omitempty"`

	// **参数解释：** 配置加速域名的源站，CDN通过加速域名将源站上需要加速的内容缓存到CDN节点，达到加速访问资源的效果 **约束限制：** 不涉及
	Sources *[]SourcesConfig `json:"sources,omitempty"`

	// **参数解释：** 指定CDN回源时的请求协议 **约束限制：** 不涉及 **取值范围：** - follow: 协议跟随 - http: http协议 - https: https协议  **默认取值：** http: http协议
	OriginProtocol *string `json:"origin_protocol,omitempty"`

	// **参数解释：** 开启回源跟随，当源站地址因业务需求做了301/302 重定向，CDN节点会先跳转到301/302对应地址获取资源，缓存后再返回给用户 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	OriginFollow302Status *string `json:"origin_follow302_status,omitempty"`

	// **参数解释：** 通过不同参数控制源站资源在CDN节点的缓存时长 **约束限制：** 必须存在一条所有文件类型的规则
	CacheRules *[]CacheRules `json:"cache_rules,omitempty"`

	IpFilter *IpFilter `json:"ip_filter,omitempty"`

	Referer *RefererConfig `json:"referer,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	Compress *Compress `json:"compress,omitempty"`

	CacheUrlParameterFilter *CacheUrlParameterFilter `json:"cache_url_parameter_filter,omitempty"`

	// **参数解释：** 配置客户端用IPv6协议访问CDN节点 **约束限制：** 不涉及 **取值范围：** - 1: 开启 - 0: 关闭  **默认取值：** 不涉及
	Ipv6Accelerate *int32 `json:"ipv6_accelerate,omitempty"`

	// **参数解释：** 将源站返回的错误状态码缓存到CDN节点，用户再次请求时由CDN直接响应给用户错误码 **约束限制：** 不涉及
	ErrorCodeCache *[]ErrorCodeCache `json:"error_code_cache,omitempty"`

	// **参数解释：** Range回源，开启后，源站在收到CDN节点回源请求时，根据HTTP请求头中的Range信息返回指定范围的数据给CDN节点 **约束限制：** 开启Range回源的前提是您的源站支持Range请求，即HTTP请求头中包含Range字段，否则可能导致回源失败 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	OriginRangeStatus *string `json:"origin_range_status,omitempty"`

	UserAgentFilter *UserAgentFilter `json:"user_agent_filter,omitempty"`

	// **参数解释：** 回源URL改写，解决回源请求URL与源站URL不匹配的问题 **约束限制：** 最多配置20条
	OriginRequestUrlRewrite *[]OriginRequestUrlRewrite `json:"origin_request_url_rewrite,omitempty"`

	// **参数解释：** 高级回源，实现根据不同的资源类型或路径回源到不同源站 **约束限制：** 最多配置20条
	FlexibleOrigin *[]FlexibleOrigins `json:"flexible_origin,omitempty"`

	// **参数解释：** 回源是否校验ETag，开启此项配置后，当用户再次请求节点已缓存资源时，会将ETag的值加入到判断资源是否过期的逻辑 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	SliceEtagStatus *string `json:"slice_etag_status,omitempty"`

	// **参数解释：** 自定义设置CDN节点回源时的超时时间，请求超时即表示此次回源失败 **约束限制：** 不涉及 **取值范围：** 5-300，单位：秒 **默认取值：** 不涉及
	OriginReceiveTimeout *int32 `json:"origin_receive_timeout,omitempty"`

	RemoteAuth *CommonRemoteAuth `json:"remote_auth,omitempty"`

	Websocket *WebSocketSeek `json:"websocket,omitempty"`

	VideoSeek *VideoSeek `json:"video_seek,omitempty"`

	// **参数解释：** 请求限速，将用户请求速度限制在指定范围内，一定程度上减少突发高带宽风险，节省成本 **约束限制：** 不涉及
	RequestLimitRules *[]RequestLimitRules `json:"request_limit_rules,omitempty"`

	IpFrequencyLimit *IpFrequencyLimit `json:"ip_frequency_limit,omitempty"`

	Hsts *Hsts `json:"hsts,omitempty"`

	Quic *Quic `json:"quic,omitempty"`

	// **参数解释：** 自定义错误页面，当访问页面出错时，会给客户端返回指定的页面，优化用户体验 **约束限制：** 不涉及
	ErrorCodeRedirectRules *[]ErrorCodeRedirectRules `json:"error_code_redirect_rules,omitempty"`

	Sni *Sni `json:"sni,omitempty"`

	// **参数解释：** 访问URL重写，当用户请求URL与CDN节点缓存资源的URL不一致时，可以通过访问URL重写功能重定向到目标URL **约束限制：** 不涉及
	RequestUrlRewrite *[]RequestUrlRewrite `json:"request_url_rewrite,omitempty"`

	// **参数解释：** 浏览器缓存过期时间，当终端用户请求资源时，如果浏览器有缓存，直接返回给用户 **约束限制：** 不涉及
	BrowserCacheRules *[]BrowserCacheRules `json:"browser_cache_rules,omitempty"`

	AccessAreaFilter *[]AccessAreaFilter `json:"access_area_filter,omitempty"`

	ClientCert *ClientCert `json:"client_cert,omitempty"`

	HttpsTlsVersion *HttpsTlsVersion `json:"https_tls_version,omitempty"`

	// **参数解释：** 设置用量封顶阈值，当实际用量大于阈值时停用域名，有效预防流量盗刷或恶意攻击带来的高额账单。  > 由于监控数据存在时延，域名将在用量达到阈值后的10分钟左右被停用  **约束限制：** 不涉及
	FlowLimitStrategy *[]FlowLimitStrategy `json:"flow_limit_strategy,omitempty"`

	BandwidthAlarmStrategy *BandwidthAlarmStrategy `json:"bandwidth_alarm_strategy,omitempty"`
}

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
