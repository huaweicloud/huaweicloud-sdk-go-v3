package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails 为IAM委托授予高风险系统身份策略分析详细结果。
type IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails struct {

	// 委托的唯一标识符（ID）。
	AgencyId string `json:"agency_id"`

	// 策略名称。
	PolicyName string `json:"policy_name"`
}

func (o IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails struct{}"
	}

	return strings.Join([]string{"IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails", string(data)}, " ")
}
