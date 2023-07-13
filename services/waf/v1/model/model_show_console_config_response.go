package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsoleConfigResponse Response Object
type ShowConsoleConfigResponse struct {

	// 是否支持EPS，false：不支持；true：支持
	Eps *bool `json:"eps,omitempty"`

	// 是否支持的TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，false：不支持；true：支持
	Tls *bool `json:"tls,omitempty"`

	// 是否支持IPV6，false：不支持；true：支持
	Ipv6 *bool `json:"ipv6,omitempty"`

	// 是否支持告警，false：不支持；true：支持
	Alert *bool `json:"alert,omitempty"`

	// 是否支持精准防护，false：不支持；true：支持
	Custom *bool `json:"custom,omitempty"`

	// 是否支持ELB模式，false：不支持；true：支持
	ElbMode *bool `json:"elb_mode,omitempty"`

	// 是否支持LTS全量日志，false：不支持；true：支持
	EventLts *bool `json:"event_lts,omitempty"`

	// 是否支持多DNS解析，false：不支持；true：支持
	MultiDns *bool `json:"multi_dns,omitempty"`

	// 是否支持搜索IP，false：不支持；true：支持
	SearchIp *bool `json:"search_ip,omitempty"`

	// 是否支持CC增强，false：不支持；true：支持
	CcEnhance *bool `json:"cc_enhance,omitempty"`

	// 是否支持cname切换，false：不支持；true：支持
	CnameSwitch *bool `json:"cname_switch,omitempty"`

	// 是否支持自定义拦截页面，false：不支持，true：支持
	CustomBlock *bool `json:"custom_block,omitempty"`

	// 是否支持误报屏蔽，false：不支持；true：支持
	AdvancedIgnore *bool `json:"advanced_ignore,omitempty"`

	// 是否支持js反爬虫，false：不支持；true：支持
	JsCrawlerEnable *bool `json:"js_crawler_enable,omitempty"`

	// 是否支持web基础防护深度检测，false：不支持；true：支持
	DeepDecodeEnable *bool `json:"deep_decode_enable,omitempty"`

	// 是否支持安全总览带宽统计，false：不支持；true：支持
	OverviewBandwidth *bool `json:"overview_bandwidth,omitempty"`

	// 是否支持使用旧cname解析，false：不支持；true：支持
	ProxyUseOldcname *bool `json:"proxy_use_oldcname,omitempty"`

	// 是否支持检查所有的header，false：不支持；true：支持
	CheckAllHeadersEnable *bool `json:"check_all_headers_enable,omitempty"`

	// 是否支持地理位置访问控制，false：不支持；true：支持
	GeoipEnable *bool `json:"geoip_enable,omitempty"`

	// 是否支持域名访问负载均衡配置，false：不支持；true：支持
	LoadBalanceEnable *bool `json:"load_balance_enable,omitempty"`

	// 是否支持ipv6防护，false：不支持；true：支持
	Ipv6ProtectionEnable *bool `json:"ipv6_protection_enable,omitempty"`

	// 是否支持策略共享，false：不支持；true：支持
	PolicySharingEnable *bool `json:"policy_sharing_enable,omitempty"`

	// 是否支持ip地址组，false：不支持；true：支持
	IpGroup *bool `json:"ip_group,omitempty"`

	// 是否支持网站反爬虫，false：不支持；true：支持
	RobotActionEnable *bool `json:"robot_action_enable,omitempty"`

	// 是否支持http2，false：不支持；true：支持
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 是否支持超时配置，false：不支持；true：支持
	TimeoutConfigEnable *bool `json:"timeout_config_enable,omitempty"`
	HttpStatusCode      int   `json:"-"`
}

func (o ShowConsoleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigResponse", string(data)}, " ")
}
