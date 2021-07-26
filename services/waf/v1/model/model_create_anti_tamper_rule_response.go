package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAntiTamperRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 防篡改的域名

	Hostname *string `json:"hostname,omitempty"`
	// 防篡改的url

	Url *string `json:"url,omitempty"`
	// 创建规则的时间戳

	Description *string `json:"description,omitempty"`
	// 规则状态

	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAntiTamperRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAntiTamperRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiTamperRuleResponse", string(data)}, " ")
}
