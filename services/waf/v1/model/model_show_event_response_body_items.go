package model

import (
	"encoding/json"

	"strings"
)

// 攻击类型
type ShowEventResponseBodyItems struct {
	// 次数

	Time *int64 `json:"time,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 源ip

	Sip *string `json:"sip,omitempty"`
	// 域名

	Host *string `json:"host,omitempty"`
	// 攻击的url链接

	Url *string `json:"url,omitempty"`
	// 攻击类型

	Attack *string `json:"attack,omitempty"`
	// 命中的规则id

	Rule *string `json:"rule,omitempty"`
	// 命中的载荷

	Payload *string `json:"payload,omitempty"`
	// 防护动作

	Action *string `json:"action,omitempty"`
	// 时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o ShowEventResponseBodyItems) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEventResponseBodyItems struct{}"
	}

	return strings.Join([]string{"ShowEventResponseBodyItems", string(data)}, " ")
}
