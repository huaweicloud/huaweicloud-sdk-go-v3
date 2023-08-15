package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimplePremiumWafHost 独享模式域名信息列表
type SimplePremiumWafHost struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region *string `json:"region,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 防护域名初始绑定的防护策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 域名防护状态：  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 域名id，和id的值是一样的，属于冗余字段
	Hostid *string `json:"hostid,omitempty"`
}

func (o SimplePremiumWafHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimplePremiumWafHost struct{}"
	}

	return strings.Join([]string{"SimplePremiumWafHost", string(data)}, " ")
}
