package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateConfigs **参数解释：** 域名模板配置 **约束限制：** 不涉及
type TemplateConfigs struct {
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	CacheRules *[]CacheRules `json:"cache_rules,omitempty"`

	// **参数解释：** 开启回源跟随，当源站地址因业务需求做了301/302 重定向，CDN节点会先跳转到301/302对应地址获取资源，缓存后再返回给用户 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭 **默认取值：** 不涉及
	OriginFollow302Status *string `json:"origin_follow302_status,omitempty"`

	Compress *Compress `json:"compress,omitempty"`

	// **参数解释：** Range回源，开启后，源站在收到CDN节点回源请求时，根据HTTP请求头中的Range信息返回指定范围的数据给CDN节点 **约束限制：** 开启Range回源的前提是您的源站支持Range请求，即HTTP请求头中包含Range字段，否则可能导致回源失败 **取值范围：** - on: 开启 - off: 关闭 **默认取值：** 不涉及
	OriginRangeStatus *string `json:"origin_range_status,omitempty"`

	IpFilter *IpFilter `json:"ip_filter,omitempty"`

	Referer *RefererConfig `json:"referer,omitempty"`

	UserAgentFilter *UserAgentFilter `json:"user_agent_filter,omitempty"`

	// **参数解释：** 设置用量封顶阈值，当实际用量大于阈值时停用域名，有效预防流量盗刷或恶意攻击带来的高额账单。  > 由于监控数据存在时延，域名将在用量达到阈值后的10分钟左右被停用  **约束限制：** 不涉及
	FlowLimitStrategy *[]FlowLimitStrategy `json:"flow_limit_strategy,omitempty"`
}

func (o TemplateConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateConfigs struct{}"
	}

	return strings.Join([]string{"TemplateConfigs", string(data)}, " ")
}
