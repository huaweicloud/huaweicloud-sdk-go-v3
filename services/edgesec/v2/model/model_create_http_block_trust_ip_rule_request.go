package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpBlockTrustIpRuleRequest Request Object
type CreateHttpBlockTrustIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpBlockTrustIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpBlockTrustIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpBlockTrustIpRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpBlockTrustIpRuleRequest", string(data)}, " ")
}
