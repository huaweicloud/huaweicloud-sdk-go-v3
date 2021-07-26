package model

import (
	"encoding/json"

	"strings"
)

// 选项
type PolicyOption struct {
	// 基础防护是否开启

	Webattack *bool `json:"webattack,omitempty"`
	// 常规检测是否开启

	Common *bool `json:"common,omitempty"`
	// 反爬虫是否开启

	Crawler *bool `json:"crawler,omitempty"`
	// 搜索engine是否开启

	CrawlerEngine *bool `json:"crawler_engine,omitempty"`
	// 反爬虫检测是否开启

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
	// 黑白名单防护是否开启

	Whiteblackip *bool `json:"whiteblackip,omitempty"`
	// 误报屏蔽是否开启

	Ignore *bool `json:"ignore,omitempty"`
	// 隐私屏蔽是否开启

	Privacy *bool `json:"privacy,omitempty"`
	// 网页防篡改规则是否开启

	Antitamper *bool `json:"antitamper,omitempty"`
}

func (o PolicyOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PolicyOption struct{}"
	}

	return strings.Join([]string{"PolicyOption", string(data)}, " ")
}
