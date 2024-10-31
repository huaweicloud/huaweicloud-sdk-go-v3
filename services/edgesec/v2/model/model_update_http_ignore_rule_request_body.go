package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpIgnoreRuleRequestBody struct {

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 误报路径
	Url *string `json:"url,omitempty"`

	// 需要屏蔽的规则，可屏蔽一个或者多个，屏蔽多个时使用半角符;分隔   - 当需要屏蔽某一条内置规则时，该参数值为该内置规则id,可以在Web应用防火墙控制台的防护策略->策略名称->Web基础防护的高级设置->防护规则中查询；也可以在防护事件的事件详情中查询内置规则id   - 当需要屏蔽web基础防护某一类规则时，该参数值为需要屏蔽的web基础防护某一类规则名。其中，xss：xss攻击；webshell：网站木马；vuln：其他类型攻击；sqli：sql注入攻击；robot：恶意爬虫；rfi：远程文件包含；lfi：本地文件包含；cmdi：命令注入攻击   - 当需要屏蔽Web基础防护模块，该参数值为：all   - 当需要屏蔽规则为所有检测模块时，该参数值为：bypass
	Rule string `json:"rule"`

	// 误报屏蔽模式，默认为0即旧模式
	Mode int32 `json:"mode"`

	// 防护域名或防护网站，数组长度为0时，代表规则对全部域名或防护网站生效
	Domains []string `json:"domains"`

	// 屏蔽规则url类型（前缀：prefix，等于：equal）
	UrlLogic *string `json:"url_logic,omitempty"`

	Advanced *HttpIgnoreRuleCondition `json:"advanced,omitempty"`

	// 命中条件
	Conditions []HttpIgnoreRuleCondition `json:"conditions"`
}

func (o UpdateHttpIgnoreRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpIgnoreRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpIgnoreRuleRequestBody", string(data)}, " ")
}
