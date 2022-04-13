package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConsoleConfigResponse struct {
	// 支持EPS

	Eps *bool `json:"eps,omitempty"`
	// 支持最低的TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本

	Tls *bool `json:"tls,omitempty"`
	// 支持IPV6

	Ipv6 *bool `json:"ipv6,omitempty"`
	// 支持告警

	Alert *bool `json:"alert,omitempty"`
	// 支持精准防护

	Custom *bool `json:"custom,omitempty"`
	// 支持ELB模式

	ElbMode *bool `json:"elb_mode,omitempty"`
	// 支持LTS全量日志

	EventLts *bool `json:"event_lts,omitempty"`
	// 支持多DNS解析

	MultiDns *bool `json:"multi_dns,omitempty"`
	// 支持搜索IP

	SearchIp *bool `json:"search_ip,omitempty"`
	// 支持CC增强

	CcEnhance *bool `json:"cc_enhance,omitempty"`
	// 支持cname切换

	CnameSwitch *bool `json:"cname_switch,omitempty"`
	// 支持精准拦截

	CustomBlock *bool `json:"custom_block,omitempty"`
	// 支持误报屏蔽

	AdvancedIgnore *bool `json:"advanced_ignore,omitempty"`
	// 支持js反爬虫

	JsCrawlerEnable *bool `json:"js_crawler_enable,omitempty"`
	// 支持深度解析

	DeepDecodeEnable *bool `json:"deep_decode_enable,omitempty"`
	// 支持安全总览带宽统计

	OverviewBandwidth *bool `json:"overview_bandwidth,omitempty"`
	// 支持使用旧cname解析

	ProxyUseOldcname *bool `json:"proxy_use_oldcname,omitempty"`
	// 支持检查所有的header

	CheckAllHeadersEnable *bool `json:"check_all_headers_enable,omitempty"`
	HttpStatusCode        int   `json:"-"`
}

func (o ShowConsoleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigResponse", string(data)}, " ")
}
