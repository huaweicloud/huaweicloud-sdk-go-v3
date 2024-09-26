package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpPolicyRuleStatusResponse Response Object
type UpdateHttpPolicyRuleStatusResponse struct {

	// 规则类别
	Category *string `json:"category,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则状态
	Status *int32 `json:"status,omitempty"`

	// 规则添加时间
	Timestamp      float32 `json:"timestamp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHttpPolicyRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpPolicyRuleStatusResponse", string(data)}, " ")
}
