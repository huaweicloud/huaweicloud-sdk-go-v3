package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIgnoreRuleResponse Response Object
type CreateIgnoreRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 需要屏蔽的规则，可屏蔽一个或者多个，屏蔽多个时使用半角符;分隔   - 当需要屏蔽某一条内置规则时，该参数值为该内置规则id,可以在Web应用防火墙控制台的防护策略->策略名称->Web基础防护的高级设置->防护规则中查询；也可以在防护事件的事件详情中查询内置规则id   - 当需要屏蔽web基础防护某一类规则时，该参数值为需要屏蔽的web基础防护某一类规则名。其中，xss：xss攻击；webshell：网站木马；vuln：其他类型攻击；sqli：sql注入攻击；robot：恶意爬虫；rfi：远程文件包含；lfi：本地文件包含；cmdi：命令注入攻击   - 当需要屏蔽Web基础防护模块，该参数值为：all   - 当需要屏蔽规则为所有检测模块时，该参数值为：bypass
	Rule *string `json:"rule,omitempty"`

	// 版本号固定值为1,代表v2版本误报屏蔽规则，v1版本仅支持兼容旧版本，不支持创建
	Mode *int32 `json:"mode,omitempty"`

	// 条件列表
	Conditions *[]Condition `json:"conditions,omitempty"`

	Advanced *IgnoreAdvanced `json:"advanced,omitempty"`

	// 防护域名或防护网站
	Domain         *[]string `json:"domain,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateIgnoreRuleResponse", string(data)}, " ")
}
