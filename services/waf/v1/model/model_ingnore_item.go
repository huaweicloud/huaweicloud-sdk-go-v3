package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IngnoreItem struct {
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
	// 防篡改规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀

	Url *string `json:"url,omitempty"`
	// 屏蔽的规则，可以是命中规则id，对应防护事件的命中规则；或者所有规则（所有规则：all）；或者攻击类型枚举（XSS攻击：xss，sqli，命令注入：cmdi，恶意爬虫：robot，本地文件包含：lfi，远程文件包含：rfi，网站木马：webshell，cc攻击：cc，精准防护：custom_custom，IP黑白名单：custom_whiteblackip，地理位置访问控制：custom_geoip，防篡改：antitamper，反爬虫：anticrawler，网站信息防泄漏：leakage，非法请求：illegal，其它类型攻击：vuln）

	Rule *string `json:"rule,omitempty"`
	// 防护的域名

	Domain *[]string `json:"domain,omitempty"`
	// url匹配逻辑（prefix：前缀匹配，equal：全等）

	UrlLogic *string `json:"url_logic,omitempty"`

	Advanced *Advance `json:"advanced,omitempty"`
}

func (o IngnoreItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IngnoreItem struct{}"
	}

	return strings.Join([]string{"IngnoreItem", string(data)}, " ")
}
