package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyRuleStatusRequestBody struct {

	// 状态（开启：1，关闭：0）
	Status *int32 `json:"status,omitempty"`
}

func (o UpdatePolicyRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusRequestBody", string(data)}, " ")
}
