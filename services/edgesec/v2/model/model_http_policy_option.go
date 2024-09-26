package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpPolicyOption 选项
type HttpPolicyOption struct {

	// 基础防护是否开启
	Webattack *bool `json:"webattack,omitempty"`

	// 常规检测是否开启
	Common *bool `json:"common,omitempty"`

	// 所有反爬虫是否开启
	BotEnable *bool `json:"bot_enable,omitempty"`

	// 特征反爬虫是否开启
	Crawler *bool `json:"crawler,omitempty"`

	// 搜索engine是否开启
	CrawlerEngine *bool `json:"crawler_engine,omitempty"`

	// 扫描器是否开启
	CrawlerScanner *bool `json:"crawler_scanner,omitempty"`

	// 脚本反爬虫是否开启
	CrawlerScript *bool `json:"crawler_script,omitempty"`

	// 其他爬虫是否开启
	CrawlerOther *bool `json:"crawler_other,omitempty"`

	// Webshell检测是否开启
	Webshell *bool `json:"webshell,omitempty"`

	// cc规则是否开启
	Cc *bool `json:"cc,omitempty"`

	// 精准防护是否开启
	Custom *bool `json:"custom,omitempty"`

	// 攻击惩罚是否开启
	FollowedAction *bool `json:"followed_action,omitempty"`

	// 黑白名单防护是否开启
	Whiteblackip *bool `json:"whiteblackip,omitempty"`

	// 地理位置规则是否开启
	Geoip *bool `json:"geoip,omitempty"`

	// 误报屏蔽是否开启
	Ignore *bool `json:"ignore,omitempty"`

	// 隐私屏蔽是否开启
	Privacy *bool `json:"privacy,omitempty"`

	// 网页防篡改规则是否开启
	Antitamper *bool `json:"antitamper,omitempty"`

	// 防敏感信息泄露规则是否开启
	Antileakage *bool `json:"antileakage,omitempty"`

	// 脚本反爬虫规则是否开启
	Anticrawler *bool `json:"anticrawler,omitempty"`

	// 三方BOT是否开启
	ThirdBotRiver *bool `json:"third_bot_river,omitempty"`
}

func (o HttpPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpPolicyOption struct{}"
	}

	return strings.Join([]string{"HttpPolicyOption", string(data)}, " ")
}
