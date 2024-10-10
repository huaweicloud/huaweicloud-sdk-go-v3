package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WafPolicyOptions struct {

	// 是否开启CC（频率控制）
	Cc *bool `json:"cc,omitempty"`

	// 是否开启精准访问防护
	Custom *bool `json:"custom,omitempty"`

	// 是否开启区域封禁防护
	Geoip *bool `json:"geoip,omitempty"`

	// 是否开启黑白名单防护
	Whiteblackip *bool `json:"whiteblackip,omitempty"`

	// 是否开启智能CC防护
	ModulexEnabled *bool `json:"modulex_enabled,omitempty"`
}

func (o WafPolicyOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafPolicyOptions struct{}"
	}

	return strings.Join([]string{"WafPolicyOptions", string(data)}, " ")
}
