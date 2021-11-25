package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCompositeHostResponse struct {
	// 域名id

	Id *string `json:"id,omitempty"`
	// 创建的云模式防护域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// cname前缀

	AccessCode *string `json:"access_code,omitempty"`
	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 接入状态

	AccessStatus *int32 `json:"access_status,omitempty"`
	// 是否开启了代理

	Proxy *bool `json:"proxy,omitempty"`
	// 创建防护域名的时间

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 套餐付费模式，目前只支持prePaid预付款模式

	PaidType *string `json:"paid_type,omitempty"`

	Flag *HostFlag `json:"flag,omitempty"`
	// 域名所属WAF模式

	WafType        *string `json:"waf_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCompositeHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompositeHostResponse struct{}"
	}

	return strings.Join([]string{"ShowCompositeHostResponse", string(data)}, " ")
}
