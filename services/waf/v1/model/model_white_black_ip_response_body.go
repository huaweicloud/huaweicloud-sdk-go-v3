package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WhiteBlackIpResponseBody struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 黑白名单规则名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 创建规则的时间戳（毫秒）
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty" xml:"status"`

	// Ip/Ip段
	Addr *string `json:"addr,omitempty" xml:"addr"`

	// 防护动作：  - 0拦截  - 1放行  - 2仅记录
	White *int32 `json:"white,omitempty" xml:"white"`

	IpGroup *IpGroup `json:"ip_group,omitempty" xml:"ip_group"`
}

func (o WhiteBlackIpResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhiteBlackIpResponseBody struct{}"
	}

	return strings.Join([]string{"WhiteBlackIpResponseBody", string(data)}, " ")
}
