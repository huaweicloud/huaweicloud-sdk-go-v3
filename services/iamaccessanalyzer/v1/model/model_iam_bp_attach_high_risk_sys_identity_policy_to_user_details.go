package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpAttachHighRiskSysIdentityPolicyToUserDetails 为IAM用户授予高风险系统身份策略分析详细结果。
type IamBpAttachHighRiskSysIdentityPolicyToUserDetails struct {

	// 用户的唯一标识符（ID）。
	UserId string `json:"user_id"`

	// 策略名称。
	PolicyName string `json:"policy_name"`
}

func (o IamBpAttachHighRiskSysIdentityPolicyToUserDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpAttachHighRiskSysIdentityPolicyToUserDetails struct{}"
	}

	return strings.Join([]string{"IamBpAttachHighRiskSysIdentityPolicyToUserDetails", string(data)}, " ")
}
