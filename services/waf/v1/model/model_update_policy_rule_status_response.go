package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePolicyRuleStatusResponse struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 规则创建时间
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 状态（开启：1，关闭：0）
	Status         *int32 `json:"status,omitempty" xml:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePolicyRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusResponse", string(data)}, " ")
}
