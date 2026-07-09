package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAntitamperRuleResponse Response Object
type DeleteAntitamperRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 防篡改的url
	Url *string `json:"url,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 防护域名
	Hostname *string `json:"hostname,omitempty"`

	// 规则状态
	Status *string `json:"status,omitempty"`

	// 规则描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAntitamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntitamperRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAntitamperRuleResponse", string(data)}, " ")
}
