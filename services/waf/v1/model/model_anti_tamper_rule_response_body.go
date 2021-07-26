package model

import (
	"encoding/json"

	"strings"
)

type AntiTamperRuleResponseBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 防篡改的域名

	Hostname *string `json:"hostname,omitempty"`
	// 防篡改的url

	Url *string `json:"url,omitempty"`
	// 创建规则的时间戳

	Description *string `json:"description,omitempty"`
}

func (o AntiTamperRuleResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AntiTamperRuleResponseBody struct{}"
	}

	return strings.Join([]string{"AntiTamperRuleResponseBody", string(data)}, " ")
}
