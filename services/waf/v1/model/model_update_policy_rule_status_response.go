package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyRuleStatusResponse Response Object
type UpdatePolicyRuleStatusResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则创建时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 状态（开启：1，关闭：0）
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePolicyRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusResponse", string(data)}, " ")
}
