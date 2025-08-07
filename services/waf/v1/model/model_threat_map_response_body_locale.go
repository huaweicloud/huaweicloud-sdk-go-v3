package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThreatMapResponseBodyLocale **参数解释：** 告警通知里，各个事件类型的描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ThreatMapResponseBodyLocale struct {

	// **参数解释：** 命令注入攻击，攻击者通过注入恶意命令来执行非授权操作 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Cmdi *string `json:"cmdi,omitempty"`

	// **参数解释：** 大模型提示词注入攻击，攻击者通过构造特殊输入来篡改AI模型的提示词 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LlmPromptInjection *string `json:"llm_prompt_injection,omitempty"`

	// **参数解释：** 网站反爬虫策略，用于阻止自动化程序非法获取网站内容 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Anticrawler *string `json:"anticrawler,omitempty"`

	// **参数解释：** 精准防护，基于特定规则的定制化安全防护策略 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomCustom *string `json:"custom_custom,omitempty"`

	// **参数解释：** 第三方BOT，来自第三方服务的自动化交互程序 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ThirdBotRiver *string `json:"third_bot_river,omitempty"`

	// **参数解释：** 恶意爬虫，用于非法获取数据或进行攻击的自动化程序 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Robot *string `json:"robot,omitempty"`

	// **参数解释：** IDC情报，基于数据中心IP地址的威胁情报 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomIdcIp *string `json:"custom_idc_ip,omitempty"`

	// **参数解释：** 目录遍历防护，防止攻击者通过目录遍历访问系统文件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AntiscanDirTraversal *string `json:"antiscan_dir_traversal,omitempty"`

	// **参数解释：** 高级BOT，具有复杂行为模式的自动化程序 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AdvancedBot *string `json:"advanced_bot,omitempty"`

	// **参数解释：** XSS攻击，跨站脚本攻击，攻击者通过注入恶意脚本获取用户信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Xss *string `json:"xss,omitempty"`

	// **参数解释：** 高频扫描封锁，对异常高频请求进行识别和拦截 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AntiscanHighFreqScan *string `json:"antiscan_high_freq_scan,omitempty"`

	// **参数解释：** 网站木马，攻击者上传的用于远程控制网站的恶意程序 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Webshell *string `json:"webshell,omitempty"`

	// **参数解释：** CC攻击，挑战型攻击，通过大量请求耗尽服务器资源 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Cc *string `json:"cc,omitempty"`

	// **参数解释：** BOT攻击，利用自动化程序进行的恶意攻击 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Botm *string `json:"botm,omitempty"`

	// **参数解释：** 非法请求，违反安全策略或业务规则的请求 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Illegal *string `json:"illegal,omitempty"`

	// **参数解释：** 大模型提示词合规检测，识别提示词中的敏感信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LlmPromptSensitive *string `json:"llm_prompt_sensitive,omitempty"`

	// **参数解释：** SQL注入，攻击者通过注入恶意SQL语句来获取或篡改数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sqli *string `json:"sqli,omitempty"`

	// **参数解释：** 本地文件包含，攻击者利用漏洞包含本地文件获取信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Lfi *string `json:"lfi,omitempty"`

	// **参数解释：** 网页防篡改，保护网站内容不被非法修改 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Antitamper *string `json:"antitamper,omitempty"`

	// **参数解释：** 地理位置访问控制，基于地理位置的访问限制策略 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomGeoip *string `json:"custom_geoip,omitempty"`

	// **参数解释：** 远程文件包含，攻击者利用漏洞包含远程文件执行恶意代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Rfi *string `json:"rfi,omitempty"`

	// **参数解释：** 其他类型攻击，未归类的安全漏洞或攻击 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Vuln *string `json:"vuln,omitempty"`

	// **参数解释：** 大模型响应合规检测，识别AI模型输出中的敏感信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LlmResponseSensitive *string `json:"llm_response_sensitive,omitempty"`

	// **参数解释：** IP黑白名单，基于IP地址的访问控制策略 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomWhiteblackip *string `json:"custom_whiteblackip,omitempty"`

	// **参数解释：** 网站信息泄露，敏感信息意外暴露的安全事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Leakage *string `json:"leakage,omitempty"`
}

func (o ThreatMapResponseBodyLocale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThreatMapResponseBodyLocale struct{}"
	}

	return strings.Join([]string{"ThreatMapResponseBodyLocale", string(data)}, " ")
}
