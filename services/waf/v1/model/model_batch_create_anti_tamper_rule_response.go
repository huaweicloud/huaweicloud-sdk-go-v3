package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateAntiTamperRuleResponse Response Object
type BatchCreateAntiTamperRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 防篡改的域名
	Hostname *string `json:"hostname,omitempty"`

	// 防篡改的url，
	Url *string `json:"url,omitempty"`

	// 创建规则的时间戳
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateAntiTamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAntiTamperRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAntiTamperRuleResponse", string(data)}, " ")
}
