package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompositeHostResponse struct {
	// 域名id

	Id *string `json:"id,omitempty"`
	// 创建的云模式防护域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// cname前缀

	AccessCode *string `json:"access_code,omitempty"`
	// 防护状态

	ProtectStatus *int32 `json:"protect_status,omitempty"`
	// 接入状态

	AccessStatus *int32 `json:"access_status,omitempty"`
	// 是否开启了代理

	Proxy *bool `json:"proxy,omitempty"`
	// 创建防护域名的时间

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 付费模式

	PaidType *string `json:"paid_type,omitempty"`

	Flag *HostFlag `json:"flag,omitempty"`
	// 域名所属WAF模式

	WafType *string `json:"waf_type,omitempty"`
}

func (o CompositeHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompositeHostResponse struct{}"
	}

	return strings.Join([]string{"CompositeHostResponse", string(data)}, " ")
}
